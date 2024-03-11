package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseExternalUserIDToPendingID struct {
	response.ResponseWork
	// Result 转换结果
	Result []ExternalUserIDToPendingIDResult `json:"result,omitempty"`
}

// ExternalUserIDToPendingIDResult 转换结果
type ExternalUserIDToPendingIDResult struct {
	// ExternalUserID 该授权企业的外部联系人ID
	ExternalUserID string `json:"external_userid,omitempty"`
	// PendingID 该微信账号还未成为企业客户时，unionid_to_external_userid接口返回的临时外部联系人ID
	PendingID string `json:"pending_id,omitempty"`
}
