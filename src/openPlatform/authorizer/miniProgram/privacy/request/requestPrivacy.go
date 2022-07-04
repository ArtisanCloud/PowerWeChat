package request

type OwnerSetting struct {
	ContactEmail         string `json:"contact_email"`
	ContactPhone         string `json:"contact_phone"`
	ContactQq            string `json:"contact_qq"`
	ContactWeixin        string `json:"contact_weixin"`
	ExtFileMediaId       string `json:"ext_file_media_id"`
	NoticeMethod         string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}

type SDK struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyText string `json:"privacy_text"`
}

type SDKPrivacyInfo struct {
	SDKName    string `json:"sdk_name"`
	SDKBizName string `json:"sdk_biz_name"`
	SDKList    []*SDK `json:"sdk_list"`
}

type Setting struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyText string `json:"privacy_text"`
}

type RequestSet struct {
	OwnerSetting       *OwnerSetting     `json:"owner_setting"`
	SettingList        []*Setting        `json:"setting_list"`
	SDKPrivacyInfoList []*SDKPrivacyInfo `json:"sdk_privacy_info_list"`
	PrivacyVer         int               `json:"privacy_ver"`
}
