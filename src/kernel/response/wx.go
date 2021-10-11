package response

type ResponseWork struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`
}

type ResponsePayment struct {
	ReturnCode string `json:"return_code,omitempty"`
	ReturnMSG  string `json:"return_msg,omitempty"`

	ResultCode string `json:"result_code,omitempty"` // 是	String(16)	SUCCESS/FAIL
	ErrCode    string `json:"err_code,omitempty"`    // 否	String(32)	SYSTEMERROR--系统错误
	ErrMSG     string `json:"errmsg,omitempty"`
	ErrCodeDes string `json:"err_code_des,omitempty"`
}

type ResponseMiniProgram struct {
	Msg     string `json:"msg,omitempty"` // 小程序直播的部分接口会把错误提示抛在msg字段
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`
}

type ResponseOfficialAccount struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`
}

const (
	RESPONSE_TYPE_RAW = "raw"
	RESPONSE_TYPE_MAP = "map"
)

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
