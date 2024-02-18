package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseUnionIDToExternalUserID struct {
	response.ResponseWork
	// ExternalUserID 该授权企业的外部联系人ID
	ExternalUserID string `json:"external_userid,omitempty"`
	// PendingID 该微信账号尚未成为企业客户时，返回的临时外部联系人ID，该ID有效期为90天，当该用户在90天内成为企业客户时，可以通过external_userid查询pending_id关联
	PendingID string `json:"pending_id,omitempty"`
}
