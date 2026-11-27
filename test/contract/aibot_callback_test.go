package contract_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot"
	"github.com/stretchr/testify/require"
)

// Generated with WeCom's WXBizJsonMsgCrypt Python 3 example (document 101039).
const botKey = "AAECAwQFBgcICQoLDA0ODxAREhMUFRYXGBkaGxwdHh8"

type callbackFixture struct{ ciphertext, signature string }

var (
	botEcho   = callbackFixture{"NW/mp8gC8Qs5WZxgQSUDGMoFIHDi2+6d9xteQLSraVA=", "da03f10d764d5e5e9369d7c00fd62a0d2cfc14c9"}
	botText   = callbackFixture{"+54Ts/Zd8I33Y7CVPAaPKikA/YqbYIT1qmpGS9SZgaGphnJ+qN25yCHMoTgz376gk9qTS5jm6PE9gPa3bDHXLsui0dMCR6HmrA8RtRjwzBQtbttFCAZfpYaBsedQfvX3+nmRiyApGSbFBx6aSF0YsYnoP1Mcjls89BUf7Y4rK+8=", "f650fc4abedfb4ba1f351f489e19539c1b20eca8"}
	botEvent  = callbackFixture{"jW8f3vGsJ+PenB57ty+YqVciMKbugn7+xrnMJ72qqrjyaUOVMQfETudLu/KTkRQzAtMbtFq6234+Duyoef/HMhZYkNHvSrnwywLc2TKr06i4Uur7qiocfpicXMUZTIAusrx9usMxknmdRw5TjKzB0SZEypnKsLkU507zes0HCmA=", "8781b1b615bd3edd6df446b93c4fa8876984e56f"}
	botStream = callbackFixture{"+XEg0yl6ElHKIhCE4ZqBEvZwSGyKGMiYK7gEu71r+8Pa8j5BItAyGRFPg2Or2DXdsTnqg6U5APj53OU69qCGLw==", "24ba001592e62ceeb6797f48693bf5f9af01500e"}
	otherBot  = callbackFixture{"KgmNari+AkUnLI7kgqxkpF+UXNPAPyqc/wXtpqc/5SXYkbdeRYBBnyASyz6CZVqKbQvl9J1FQDzF0iEq1bJxiQ==", "5bde3d4c76e18dea28387348be0b037d3f38a31e"}
)

func (f callbackFixture) call(h http.Handler, method string) *httptest.ResponseRecorder {
	q := url.Values{"msg_signature": {f.signature}, "timestamp": {"1700000000"}, "nonce": {"nonce"}}
	var body string
	if method == http.MethodGet {
		q.Set("echostr", f.ciphertext)
	} else {
		body = `{"encrypt":"` + f.ciphertext + `"}`
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(method, "/bot?"+q.Encode(), strings.NewReader(body)))
	return w
}

func TestAIBotHTTPCallback(t *testing.T) {
	const reply = `{"msgtype":"stream","stream":{"id":"s1","content":"收到","finish":true}}`
	var handled []string
	h, err := aibot.NewCallbackHandler("test-token", botKey, func(ctx context.Context, msg *aibot.CallbackMessage) (any, error) {
		handled = append(handled, msg.MsgType)
		switch msg.MsgType {
		case "text":
			require.Equal(t, "alice", msg.From.UserID)
			require.Equal(t, "你好", msg.Text.Content)
			require.Equal(t, "bot", msg.AIBotID)
			return json.RawMessage(reply), nil
		case "event":
			require.Equal(t, "enter_chat", msg.Event.EventType)
			require.Contains(t, string(msg.Raw), `"future":true`)
		case "stream":
			require.Equal(t, "s1", msg.Stream.ID)
			return json.RawMessage(reply), nil
		}
		return nil, nil
	})
	require.NoError(t, err)
	echo := botEcho.call(h, http.MethodGet)
	require.Equal(t, 200, echo.Code)
	require.Equal(t, "verify+42", echo.Body.String())

	crypt, err := kernel.NewEncryptor("", "test-token", botKey)
	require.NoError(t, err)
	for _, f := range []callbackFixture{botText, botStream} {
		w := f.call(h, http.MethodPost)
		require.Equal(t, 200, w.Code)
		require.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var envelope struct {
			Encrypt      string `json:"encrypt"`
			MsgSignature string `json:"msgsignature"`
			Timestamp    int64  `json:"timestamp"`
			Nonce        string `json:"nonce"`
		}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &envelope))
		plain, cryptErr := crypt.OpenCallback(envelope.Encrypt, envelope.MsgSignature, envelope.Nonce, strconv.FormatInt(envelope.Timestamp, 10))
		require.Nil(t, cryptErr)
		require.JSONEq(t, reply, string(plain))
	}
	event := botEvent.call(h, http.MethodPost)
	require.Equal(t, 200, event.Code)
	require.Equal(t, "success", event.Body.String())
	require.Equal(t, []string{"text", "stream", "event"}, handled)
}

func TestAIBotCallbackErrors(t *testing.T) {
	calls := 0
	callback := func(context.Context, *aibot.CallbackMessage) (any, error) {
		calls++
		return nil, errors.New("business error")
	}
	_, err := aibot.NewCallbackHandler("test-token", "invalid", callback)
	require.Error(t, err)
	h, err := aibot.NewCallbackHandler("test-token", botKey, callback)
	require.NoError(t, err)

	badSignature := callbackFixture{botText.ciphertext, "invalid"}
	require.Equal(t, 403, badSignature.call(h, http.MethodPost).Code)
	require.Equal(t, 403, otherBot.call(h, http.MethodGet).Code)
	require.Equal(t, 400, botEcho.call(h, http.MethodPost).Code) // A signed challenge is not a JSON message.
	require.Equal(t, 405, botEcho.call(h, http.MethodPut).Code)
	require.Zero(t, calls)
	require.Equal(t, 500, botEvent.call(h, http.MethodPost).Code)
	require.Equal(t, 1, calls)
}
