package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGroupChatTransfer struct {
	response.ResponseWork

	ExternalContactList []*ResponseFailedChat `json:"failed_chat_list"`
}

type ResponseFailedChat struct {
	response.ResponseWork

	ChatID  string `json:"chat_id"` //  "wrOgQhDgAAcwMTB7YmDkbeBsgT_KAAAA",
	ErrCode int    `json:"errcode"` //  90500,
	ErrMsg  string `json:"errmsg"`  //  "the owner of this chat is not resigned"

}
