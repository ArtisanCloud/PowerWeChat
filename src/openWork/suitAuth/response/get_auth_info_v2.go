package response

// GetAuthInfoV2 shares the public response type used by GetAuthInfoV2 in v3.
type GetAuthInfoV2 = GetPermanentCodeResponse

// CorpExName 企业其他认证的名称
type CorpExName struct {
	// NameList 企业其他认证的企业简称列表（不包括 corp_name）
	NameList []string `json:"name_list,omitempty"`
}

// SharedFrom 共享了应用的企业信息
type SharedFrom struct {
	// CorpID 共享了应用的企业 corpid，仅当企业互联或者上下游共享应用触发的安装时才返回
	CorpID string `json:"corpid,omitempty"`
	// ShareType 共享了途径，0 表示企业互联，1 表示上下游
	ShareType int `json:"share_type,omitempty"`
}
