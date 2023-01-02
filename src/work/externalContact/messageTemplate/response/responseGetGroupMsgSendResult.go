package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type SendResult struct {
	ExternalUserID string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int    `json:"send_time"`
}

type ResponseGetGroupMsgSendResult struct {
	response.ResponseWork

	NextCursor string `json:"next_cursor"`

	SendList []*SendResult `json:"send_list"`
}
