package request

type RequestRiskControl struct {
	AppID        string `json:"appid"`
	OpenID       string `json:"openid"`
	Scene        int    `json:"scene"`
	MobileNo     string `json:"mobile_no,omitempty"`
	BankCardNo   string `json:"bank_card_no"`
	CertNo       string `json:"cert_no"`
	ClientIP     string `json:"client_ip"`
	EmailAddress string `json:"email_address,omitempty"`
	ExtendedInfo string `json:"extended_info,omitempty"`
}
