package response

import "github.com/ArtisanCloud/powerwechat/src/kernel/response"

type ResponseAddMessageTemplate struct {
	*response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
