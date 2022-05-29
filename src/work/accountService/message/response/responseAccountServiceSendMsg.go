package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	*response.ResponseWork

	MsgID string `json:"msgid"`
}
