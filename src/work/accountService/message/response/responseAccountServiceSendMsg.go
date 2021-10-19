package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	*response.ResponseWork

	MsgID string            `json:"msgid"`
}
