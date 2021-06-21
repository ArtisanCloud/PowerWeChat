package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetExternalContact struct {
	response.ResponseWX
	ExternalContact *ResponseExternalContact `json:"external_contact_list"`
}
