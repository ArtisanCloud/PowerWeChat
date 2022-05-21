package request

type RequestShakeAroundAccountRegister struct {
	Name                  string   `json:"name"`
	PhoneNumber           string   `json:"phone_number"`
	Email                 string   `json:"email"`
	IndustryID            string   `json:"industry_id"`
	QualificationCertUrls []string `json:"qualification_cert_urls"`
	ApplyReason           string   `json:"apply_reason"`
}


// ---------------------------------------------------------


type RequestShakeAroundUser struct {
	Ticket  string `json:"ticket"`
	NeedPoi int    `json:"need_poi"`
}
