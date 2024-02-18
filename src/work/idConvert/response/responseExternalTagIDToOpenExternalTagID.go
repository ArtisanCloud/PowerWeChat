package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseExternalTagIDToOpenExternalTagID struct {
	response.ResponseWork
	// Items 客户标签转换结果
	Items []ExternalTagIDToOpenOpenExternalTagIDResult `json:"items,omitempty"`
	// InvalidExternalTagIDList 无法转换的客户标签ID列表
	InvalidExternalTagIDList []string `json:"invalid_tagid_list,omitempty"`
}

// ExternalTagIDToOpenExternalTagIDResult 客户标签转换结果
type ExternalTagIDToOpenOpenExternalTagIDResult struct {
	// ExternalTagID 企业主体下的客户标签ID
	ExternalTagID string `json:"external_tagid,omitempty"`
	// OpenExternalTagID 服务商主体下的客户标签ID，如果传入的external_tagid已经是服务商主体下的ID，则open_external_tagid与external_tagid相同。
	OpenExternalTagID string `json:"open_external_tagid,omitempty"`
}
