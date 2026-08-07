package aibot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot/response"
	"github.com/gorilla/websocket"
)

const defaultAIBotLongConnEndpoint = "wss://openws.work.weixin.qq.com"

type MessageHandler func(ctx context.Context, payload *response.ResponseLongConnection) error

type Client struct {
	App      kernel.ApplicationInterface
	Endpoint string
	Dialer   *websocket.Dialer

	connMu  sync.RWMutex
	conn    *websocket.Conn
	writeMu sync.Mutex
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	if app == nil {
		return nil, errors.New("application is nil")
	}
	endpoint := strings.TrimSpace(readConfigString(app, "aibot.long_connection.url"))
	if endpoint == "" {
		endpoint = defaultAIBotLongConnEndpoint
	}
	return &Client{
		App:      app,
		Endpoint: endpoint,
		Dialer: &websocket.Dialer{
			HandshakeTimeout: 8 * time.Second,
		},
	}, nil
}

func (comp *Client) Connect(ctx context.Context) error {
	if comp == nil {
		return errors.New("aibot client is nil")
	}
	endpoint := strings.TrimSpace(comp.Endpoint)
	if endpoint == "" {
		return errors.New("aibot endpoint is empty")
	}
	dialer := comp.Dialer
	if dialer == nil {
		dialer = &websocket.Dialer{HandshakeTimeout: 8 * time.Second}
	}
	conn, _, err := dialer.DialContext(ctx, endpoint, nil)
	if err != nil {
		return err
	}
	comp.connMu.Lock()
	if comp.conn != nil {
		_ = comp.conn.Close()
	}
	comp.conn = conn
	comp.connMu.Unlock()
	return nil
}

func (comp *Client) Close() error {
	if comp == nil {
		return nil
	}
	comp.connMu.Lock()
	defer comp.connMu.Unlock()
	if comp.conn == nil {
		return nil
	}
	err := comp.conn.Close()
	comp.conn = nil
	return err
}

func (comp *Client) Subscribe(ctx context.Context, botID, secret string) (*response.ResponseLongConnection, error) {
	if comp == nil {
		return nil, errors.New("aibot client is nil")
	}
	botID = strings.TrimSpace(botID)
	secret = strings.TrimSpace(secret)
	if botID == "" || secret == "" {
		return nil, errors.New("botID and secret are required")
	}
	if err := comp.Connect(ctx); err != nil {
		return nil, err
	}
	cmd := request.NewSubscribe(botID, secret, buildReqID())
	if err := comp.Send(ctx, cmd); err != nil {
		return nil, err
	}
	resp, err := comp.Read(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil && resp.IsError() {
		return resp, errors.New(resp.ErrMsg)
	}
	return resp, nil
}

func (comp *Client) Send(ctx context.Context, cmd *request.RequestLongConnection) error {
	if comp == nil {
		return errors.New("aibot client is nil")
	}
	if cmd == nil {
		return errors.New("command is nil")
	}
	conn, err := comp.mustConn()
	if err != nil {
		return err
	}
	if dl, ok := ctx.Deadline(); ok {
		_ = conn.SetWriteDeadline(dl)
	}
	comp.writeMu.Lock()
	defer comp.writeMu.Unlock()
	return conn.WriteJSON(cmd)
}

func (comp *Client) Read(ctx context.Context) (*response.ResponseLongConnection, error) {
	conn, err := comp.mustConn()
	if err != nil {
		return nil, err
	}
	if dl, ok := ctx.Deadline(); ok {
		_ = conn.SetReadDeadline(dl)
	}
	_, raw, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	resp := &response.ResponseLongConnection{}
	if err := json.Unmarshal(raw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (comp *Client) Heartbeat(ctx context.Context) error {
	cmd := request.NewPing(buildReqID())
	return comp.Send(ctx, cmd)
}

func (comp *Client) Listen(ctx context.Context, handler MessageHandler) error {
	if comp == nil {
		return errors.New("aibot client is nil")
	}
	if handler == nil {
		return errors.New("handler is nil")
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		message, err := comp.Read(ctx)
		if err != nil {
			return err
		}
		if message == nil {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(message.Cmd), "pong") {
			continue
		}
		if handleErr := handler(ctx, message); handleErr != nil {
			return handleErr
		}
	}
}

func (comp *Client) SubscribeAndServe(ctx context.Context, botID, secret string, handler MessageHandler) error {
	if _, err := comp.Subscribe(ctx, botID, secret); err != nil {
		return err
	}

	heartbeatTicker := time.NewTicker(30 * time.Second)
	defer heartbeatTicker.Stop()

	errCh := make(chan error, 2)
	go func() {
		errCh <- comp.Listen(ctx, handler)
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- nil
				return
			case <-heartbeatTicker.C:
				hbCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
				hbErr := comp.Heartbeat(hbCtx)
				cancel()
				if hbErr != nil {
					errCh <- hbErr
					return
				}
			}
		}
	}()

	select {
	case <-ctx.Done():
		_ = comp.Close()
		return nil
	case err := <-errCh:
		_ = comp.Close()
		return err
	}
}

func (comp *Client) SendCommand(ctx context.Context, cmd string, body map[string]interface{}) error {
	command := request.NewCommand(cmd, buildReqID(), body)
	if strings.TrimSpace(command.Cmd) == "" {
		return errors.New("command cmd is empty")
	}
	return comp.Send(ctx, command)
}

func (comp *Client) RespondWelcomeMessage(ctx context.Context, body map[string]interface{}) error {
	return comp.SendCommand(ctx, request.CmdRespondWelcome, body)
}

func (comp *Client) RespondMessage(ctx context.Context, body map[string]interface{}) error {
	return comp.SendCommand(ctx, request.CmdRespondMessage, body)
}

func (comp *Client) RespondUpdateMessage(ctx context.Context, body map[string]interface{}) error {
	return comp.SendCommand(ctx, request.CmdRespondUpdateMsg, body)
}

func (comp *Client) SendMessage(ctx context.Context, body map[string]interface{}) error {
	return comp.SendCommand(ctx, request.CmdSendMessage, body)
}

func (comp *Client) mustConn() (*websocket.Conn, error) {
	if comp == nil {
		return nil, errors.New("aibot client is nil")
	}
	comp.connMu.RLock()
	defer comp.connMu.RUnlock()
	if comp.conn == nil {
		return nil, errors.New("aibot websocket is not connected")
	}
	return comp.conn, nil
}

func buildReqID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}

func readConfigString(app kernel.ApplicationInterface, key string) string {
	if app == nil || app.GetConfig() == nil {
		return ""
	}
	return strings.TrimSpace(app.GetConfig().GetString(key, ""))
}
