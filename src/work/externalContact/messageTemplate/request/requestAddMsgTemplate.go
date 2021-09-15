package request

import (
	request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/message/request"
)

type RequestAddMsgTemplate struct {
	ExternalUserID []string               `json:"external_userid" binding:"required"`
	Sender         string                 `json:"sender" binding:"required"`
	Text           request2.TextOfMessage `json:"text" binding:"required"`
	Attachments    []interface{}          `json:"attachments"`
}
