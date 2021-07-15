package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetExternalContact struct {
	response.ResponseWork
	ExternalContact *ResponseExternalContactDetail `json:"external_contact"`
	FollowInfo      []*FollowUser                  `json:"follow_user"`
}
