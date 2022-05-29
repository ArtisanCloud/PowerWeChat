package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseAddMessageTemplate struct {
	*response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
