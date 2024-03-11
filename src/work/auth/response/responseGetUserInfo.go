package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetUserInfo struct {
	response.ResponseWork
	UserID         string `json:"userid,omitempty"`
	UserTicket     string `json:"user_ticket,omitempty"`
	OpenID         string `json:"openid,omitempty"`
	ExternalUserID string `json:"external_userid,omitempty"`
}
