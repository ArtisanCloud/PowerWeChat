package contract_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot/response"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestAIBotReplyCorrelation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	upgrader := websocket.Upgrader{}
	written := make(chan request.RequestLongConnection, 5)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		defer close(written)
		deadline, _ := ctx.Deadline()
		_ = conn.SetReadDeadline(deadline)
		var sub request.RequestLongConnection
		if conn.ReadJSON(&sub) != nil {
			return
		}
		frames := []any{
			map[string]any{"headers": sub.Headers, "errcode": 0},
			request.NewCommand(request.CmdMessageCallback, "msg-1", nil),
			request.NewCommand(request.CmdEventCallback, "event-2", nil),
			request.NewCommand(request.CmdMessageCallback, "direct-3", nil),
		}
		for _, frame := range frames {
			if conn.WriteJSON(frame) != nil {
				return
			}
		}
		for i := 0; i < cap(written); i++ {
			var frame request.RequestLongConnection
			if conn.ReadJSON(&frame) != nil {
				return
			}
			written <- frame
		}
	}))
	defer server.Close()
	bot := &aibot.Client{Endpoint: "ws" + strings.TrimPrefix(server.URL, "http")}
	defer bot.Close()
	_, err := bot.Subscribe(ctx, "bot", "secret")
	require.NoError(t, err)
	require.ErrorContains(t, bot.RespondMessage(ctx, nil), "req_id")

	var first context.Context
	stop := errors.New("handled both callbacks")
	err = bot.Listen(ctx, func(handlerCtx context.Context, frame *response.ResponseLongConnection) error {
		if frame.Cmd == request.CmdMessageCallback {
			first = handlerCtx
			return nil
		}
		require.NoError(t, bot.RespondWelcomeMessage(handlerCtx, nil))
		require.NoError(t, bot.RespondMessage(first, nil))
		require.NoError(t, bot.RespondUpdateMessage(handlerCtx, nil))
		require.NoError(t, bot.SendMessage(handlerCtx, nil))
		return stop
	})
	require.ErrorIs(t, err, stop)
	frame, err := bot.Read(ctx)
	require.NoError(t, err)
	require.NoError(t, bot.RespondMessage(aibot.WithReplyReqID(ctx, frame.Headers.ReqID), nil))

	var cmds, ids []string
	for sent := range written {
		require.NotNil(t, sent.Headers)
		cmds = append(cmds, sent.Cmd)
		ids = append(ids, sent.Headers.ReqID)
	}
	require.Equal(t, []string{"aibot_respond_welcome_msg", "aibot_respond_msg", "aibot_respond_update_msg", "aibot_send_msg", "aibot_respond_msg"}, cmds)
	require.Equal(t, []string{"event-2", "msg-1", "event-2"}, ids[:3])
	require.NotEmpty(t, ids[3])
	require.NotEqual(t, "event-2", ids[3])
	require.Equal(t, "direct-3", ids[4])
}
