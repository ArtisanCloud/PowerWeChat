package request

import "strings"

const (
	CmdSubscribe        = "aibot_subscribe"
	CmdPing             = "ping"
	CmdRespondWelcome   = "aibot_respond_welcome_msg"
	CmdRespondMessage   = "aibot_respond_msg"
	CmdRespondUpdateMsg = "aibot_respond_update_msg"
	CmdSendMessage      = "aibot_send_msg"
)

type LongConnectionHeaders struct {
	ReqID string `json:"req_id,omitempty"`
}

type RequestLongConnection struct {
	Cmd     string                 `json:"cmd"`
	Headers *LongConnectionHeaders `json:"headers,omitempty"`
	Body    map[string]interface{} `json:"body,omitempty"`
}

func NewSubscribe(botID, secret, reqID string) *RequestLongConnection {
	return &RequestLongConnection{
		Cmd: CmdSubscribe,
		Headers: &LongConnectionHeaders{
			ReqID: reqID,
		},
		Body: map[string]interface{}{
			"bot_id": botID,
			"secret": secret,
		},
	}
}

func NewPing(reqID string) *RequestLongConnection {
	return &RequestLongConnection{
		Cmd: CmdPing,
		Headers: &LongConnectionHeaders{
			ReqID: reqID,
		},
	}
}

func NewCommand(cmd, reqID string, body map[string]interface{}) *RequestLongConnection {
	req := &RequestLongConnection{
		Cmd: strings.TrimSpace(cmd),
	}
	if reqID != "" {
		req.Headers = &LongConnectionHeaders{
			ReqID: reqID,
		}
	}
	if len(body) > 0 {
		req.Body = body
	}
	return req
}
