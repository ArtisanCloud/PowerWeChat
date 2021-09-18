package response

type ResponseWork struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

type ResponsePayment struct {
	ReturnCode string `json:"return_code"`
	ReturnMSG  string `json:"return_msg"`

	ResultCode string `json:"result_code"` // 是	String(16)	SUCCESS/FAIL
	ErrCode    string `json:"err_code"`    // 否	String(32)	SYSTEMERROR--系统错误
	ErrMSG     string `json:"errmsg"`
	ErrCodeDes string `json:"err_code_des"`
}

type ResponseMiniProgram struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`

	ResultCode  string `json:"resultcode"`
	ResultMSG  string `json:"resultmsg"`

}

type ResponseOfficialAccount struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`

	ResultCode  string `json:"resultcode"`
	ResultMSG  string `json:"resultmsg"`
}

type ResponseOfficialAccount struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`

	ResultCode  string `json:"resultcode"`
	ResultMSG  string `json:"resultmsg"`
}

type ResponseOfficialAccount struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`

	ResultCode  string `json:"resultcode"`
	ResultMSG  string `json:"resultmsg"`
}

//
//func (res *ResponseWork) GetBody() *http.ResponseWriter {
//	return nil
//}
//func (res *ResponseWork) GetHeaders() *http.ResponseWriter {
//	return nil
//}
//
//func (res *ResponseWork) GetStatusCode() int {
//	return 200
//}
