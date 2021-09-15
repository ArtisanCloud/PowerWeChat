package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	*response.ResponseWork

	MsgID string            `json:"msgid"`
}
