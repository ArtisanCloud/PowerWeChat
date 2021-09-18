package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseAddMessageTemplate struct {
	*response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
