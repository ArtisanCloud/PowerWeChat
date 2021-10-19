package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	*response.ResponseWork

	MsgID string            `json:"msgid"`
}
