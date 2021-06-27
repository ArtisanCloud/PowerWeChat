package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetExternalContact struct {
	response.ResponseWX
	ExternalContact *ResponseExternalContactDetail `json:"external_contact"`
	FollowInfo      []*FollowUser                  `json:"follow_user"`
}
