package response

type SubmerchantSubmit struct {
	Info *SubmerchantInfo `json:"info"`
}

type SubmerchantInfo struct {
	MerchantID          int    `json:"merchant_id"`
	AppID               string `json:"app_id"`
	CreateTime          int    `json:"create_time"`
	UpdateTime          int    `json:"update_time"`
	BrandName           string `json:"brand_name"`
	LogoURL             string `json:"logo_url"`
	Status              string `json:"status"`
	BeginTime           int    `json:"begin_time"`
	EndTime             int    `json:"end_time"`
	PrimaryCategoryID   int    `json:"primary_category_id"`
	SecondaryCategoryID int    `json:"secondary_category_id"`
}
type ProtocolCategory struct {
	Category []struct {
		PrimaryCategoryID int    `json:"primary_category_id"`
		CategoryName      string `json:"category_name"`
		SecondaryCategory []struct {
			SecondaryCategoryID     int      `json:"secondary_category_id"`
			CategoryName            string   `json:"category_name"`
			NeedQualificationStuffs []string `json:"need_qualification_stuffs"`
			CanChoosePrepaidCard    int      `json:"can_choose_prepaid_card"`
			CanChoosePaymentCard    int      `json:"can_choose_payment_card"`
		} `json:"secondary_category"`
	} `json:"category"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
