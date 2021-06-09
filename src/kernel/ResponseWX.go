package kernel

type ResponseWX struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}
