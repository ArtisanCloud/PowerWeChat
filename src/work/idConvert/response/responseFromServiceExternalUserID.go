package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseFromServiceExternalUserID struct {
	response.ResponseWork
	// ExternalUserID 转换后的自建应用的external_userid
	ExternalUserID string `json:"external_userid,omitempty"`
}
