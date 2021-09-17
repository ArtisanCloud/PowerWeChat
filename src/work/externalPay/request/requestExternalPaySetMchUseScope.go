package request

type RequestExternalPaySetMchUseScope struct {
	MchID         string `json:"mch_id"`
	AllowUseScope string `json:"allow_use_scope"`
}
