package response

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ResponseDeviceBind struct {
	BaseResp *BaseResp `json:"base_resp"`
}


