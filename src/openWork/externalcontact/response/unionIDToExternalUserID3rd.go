package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type UnionIDToExternalUserID3rd struct {
	response.ResponseWork
	// ExternalUserIDInfo 该unionid对应的外部联系人信息
	ExternalUserIDInfo []ExternalUserIDInfo `json:"external_userid_info,omitempty"`
}

// ExternalUserIDInfo 该unionid对应的外部联系人信息
type ExternalUserIDInfo struct {
	// CorpID 所属企业id
	CorpID string `json:"corpid,omitempty"`
	// ExternalUserID 外部联系人id
	ExternalUserID string `json:"external_userid,omitempty"`
}

type UnionIDToExternalUserID struct {
	response.ResponseWork
	// ExternalUserID 该企业的外部联系人ID
	ExternalUserID string `json:"external_userid,omitempty"`
}
