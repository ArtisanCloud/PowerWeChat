package request

type SubMerchantSubmit struct {
	Info *SubMerchantInfo `json:"info"`
}

type SubMerchantInfo struct {
	BrandName           string `json:"brand_name"`
	AppID               string `json:"app_id,omitempty"`
	LogoURL             string `json:"logo_url"`
	Protocol            string `json:"protocol"`
	EndTime             uint   `json:"end_time"`
	PrimaryCategoryID   int    `json:"primary_category_id"`
	SecondaryCategoryID int    `json:"secondary_category_id"`
	AgreementMediaID    string `json:"agreement_media_id,omitempty"`
	OperatorMediaID     string `json:"operator_media_id,omitempty"`
}

type SubMerchantID struct {
	SubMerchantID string `json:"merchant_id"`
}
