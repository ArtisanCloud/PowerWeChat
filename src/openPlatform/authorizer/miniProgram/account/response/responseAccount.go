package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type WXVerifyInfo struct {
	QualificationVerify bool `json:"qualification_verify"`
	NamingVerify        bool `json:"naming_verify"`
}

type SignatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type HeadImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type NickNameInfo struct {
	NickName        string `json:"nickname"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type ResponseGetBasicInfo struct {
	response.ResponseOpenPlatform

	AppID             string         `json:"appid"`
	AccountType       int            `json:"account_type"`
	PrincipalType     int            `json:"principal_type"`
	PrincipalName     string         `json:"principal_name"`
	RealNameStatus    int            `json:"realname_status"`
	WXVerifyInfo      *WXVerifyInfo  `json:"wx_verify_info"`
	SignatureInfo     *SignatureInfo `json:"signature_info"`
	HeadImageInfo     *HeadImageInfo `json:"head_image_info"`
	NickName          string         `json:"nickname"`
	RegisteredCountry int            `json:"registered_country"`
	NickNameInfo      *NickNameInfo  `json:"nickname_info"`
	Credential        string         `json:"credential"`
	CustomerType      int            `json:"customer_type"`
}
