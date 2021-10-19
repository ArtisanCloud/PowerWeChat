package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	*response.ResponseWork

	MsgID string            `json:"msgid"`
}
