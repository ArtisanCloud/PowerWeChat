// Package mobileApp implements server-side WeChat login for native mobile apps.
package mobileApp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ArtisanCloud/PowerSocialite/v3/src/providers"
)

type UserConfig struct {
	AppID     string
	Secret    string
	Transport http.RoundTripper
}

type Client struct {
	appID  string
	secret string
	http   *http.Client
}

type AccessToken = providers.ResAccessToken
type UserInfo = providers.UserInfo

// NewMobileApp creates a client using the mobile application's own AppID and AppSecret.
func NewMobileApp(config *UserConfig) (*Client, error) {
	if config == nil || config.AppID == "" || config.Secret == "" {
		return nil, errors.New("mobileApp: AppID and Secret are required")
	}
	return &Client{
		appID:  config.AppID,
		secret: config.Secret,
		http: &http.Client{
			Transport: config.Transport,
			Timeout:   5 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}, nil
}

// ExchangeCode exchanges the authorization code returned by the native OpenSDK.
func (c *Client) ExchangeCode(ctx context.Context, code string) (*AccessToken, error) {
	if code == "" {
		return nil, errors.New("mobileApp: authorization code is required")
	}
	params := url.Values{
		"appid":      {c.appID},
		"secret":     {c.secret},
		"code":       {code},
		"grant_type": {"authorization_code"},
	}
	var token AccessToken
	if err := c.get(ctx, "/sns/oauth2/access_token", params, &token, &token.CommonError); err != nil {
		return nil, err
	}
	if token.OpenID == "" || token.AccessToken == "" {
		return nil, errors.New("mobileApp: token response is missing openid or access_token")
	}
	return &token, nil
}

// GetUserInfo reads the profile authorized by the user with scope snsapi_userinfo.
func (c *Client) GetUserInfo(ctx context.Context, accessToken, openID string) (*UserInfo, error) {
	if accessToken == "" || openID == "" {
		return nil, errors.New("mobileApp: access token and openid are required")
	}
	params := url.Values{"access_token": {accessToken}, "openid": {openID}}
	var user UserInfo
	if err := c.get(ctx, "/sns/userinfo", params, &user, &user.CommonError); err != nil {
		return nil, err
	}
	if user.OpenID != openID {
		return nil, errors.New("mobileApp: profile response does not match requested openid")
	}
	return &user, nil
}

func (c *Client) get(ctx context.Context, path string, params url.Values, result any, apiError *providers.CommonError) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.weixin.qq.com"+path, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = params.Encode()
	resp, err := c.http.Do(req)
	if err != nil {
		// net/http includes the URL and its credentials in transport errors.
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			err = urlErr.Err
		}
		return fmt.Errorf("mobileApp: GET %s: %w", path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("mobileApp: GET %s: HTTP %d", path, resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("mobileApp: decode %s: %w", path, err)
	}
	if apiError.ErrCode != 0 {
		return fmt.Errorf("mobileApp: %s: errcode %d: %s", path, apiError.ErrCode, apiError.ErrMsg)
	}
	return nil
}
