package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	*response.ResponseWork

	MsgID string            `json:"msgid"`
}
