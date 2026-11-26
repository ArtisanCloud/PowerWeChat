package contract_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/mobileApp"
	"github.com/stretchr/testify/require"
)

type mobileTransport func(*http.Request) (*http.Response, error)

func (f mobileTransport) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func mobileClient(t *testing.T, transport mobileTransport) *mobileApp.Client {
	t.Helper()
	app, err := mobileApp.NewMobileApp(&mobileApp.UserConfig{
		AppID: "mobile-app", Secret: "test-secret", Transport: transport,
	})
	require.NoError(t, err)
	return app
}

func mobileReply(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body))}
}

func TestMobileAppLogin(t *testing.T) {
	const code = "code+with&reserved=chars"
	var paths []string
	app := mobileClient(t, func(r *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "https", r.URL.Scheme)
		require.Equal(t, "api.weixin.qq.com", r.URL.Host)
		paths = append(paths, r.URL.Path)
		switch r.URL.Path {
		case "/sns/oauth2/access_token":
			require.Equal(t, url.Values{
				"appid": {"mobile-app"}, "secret": {"test-secret"}, "code": {code}, "grant_type": {"authorization_code"},
			}, r.URL.Query())
			return mobileReply(200, `{"access_token":"token","expires_in":7200,"refresh_token":"refresh","openid":"open","scope":"snsapi_userinfo","unionid":"union"}`), nil
		case "/sns/userinfo":
			require.Equal(t, url.Values{"access_token": {"token"}, "openid": {"open"}}, r.URL.Query())
			return mobileReply(200, `{"openid":"open","nickname":"测试用户"}`), nil
		default:
			t.Fatalf("unexpected request %s", r.URL.Path)
			return nil, nil
		}
	})
	token, err := app.ExchangeCode(context.Background(), code)
	require.NoError(t, err)
	require.Equal(t, []string{"/sns/oauth2/access_token"}, paths)
	require.Equal(t, "open", token.OpenID)
	require.Equal(t, "union", token.UnionID)
	require.Equal(t, "refresh", token.RefreshToken)
	require.EqualValues(t, 7200, token.ExpiresIn)
	user, err := app.GetUserInfo(context.Background(), token.AccessToken, token.OpenID)
	require.NoError(t, err)
	require.Equal(t, []string{"/sns/oauth2/access_token", "/sns/userinfo"}, paths)
	require.Equal(t, "测试用户", user.Nickname)
}

func TestMobileAppLoginErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
		want   string
	}{
		{"used code", 200, `{"errcode":40163,"errmsg":"code been used"}`, "40163"},
		{"missing openid", 200, `{"access_token":"token"}`, "openid"},
		{"missing token", 200, `{"openid":"open"}`, "access_token"},
		{"invalid JSON", 200, "{", "decode"},
		{"HTTP failure", 503, "", "HTTP 503"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			app := mobileClient(t, func(*http.Request) (*http.Response, error) {
				return mobileReply(tc.status, tc.body), nil
			})
			token, err := app.ExchangeCode(context.Background(), "code")
			require.ErrorContains(t, err, tc.want)
			require.Nil(t, token)
		})
	}
	t.Run("profile refused", func(t *testing.T) {
		app := mobileClient(t, func(*http.Request) (*http.Response, error) {
			return mobileReply(200, `{"errcode":40014,"errmsg":"invalid access token"}`), nil
		})
		user, err := app.GetUserInfo(context.Background(), "token", "open")
		require.ErrorContains(t, err, "40014")
		require.Nil(t, user)
	})
	t.Run("cancelled request", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		app := mobileClient(t, func(r *http.Request) (*http.Response, error) {
			cancel()
			return nil, r.Context().Err()
		})
		_, err := app.ExchangeCode(ctx, "private-code")
		require.ErrorIs(t, err, context.Canceled)
		require.NotContains(t, err.Error(), "test-secret")
		require.NotContains(t, err.Error(), "private-code")
	})
	t.Run("invalid input", func(t *testing.T) {
		_, err := mobileApp.NewMobileApp(&mobileApp.UserConfig{AppID: "mobile-app"})
		require.Error(t, err)
		app := mobileClient(t, func(*http.Request) (*http.Response, error) {
			t.Error("unexpected network request")
			return nil, errors.New("unexpected network request")
		})
		_, err = app.ExchangeCode(context.Background(), "")
		require.Error(t, err)
		_, err = app.GetUserInfo(context.Background(), "", "open")
		require.Error(t, err)
	})
}
