package response

type ResponseBase struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseWork struct {
	ResponseBase

	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`
}

type ResponsePayment struct {
	ResponseBase

	ReturnCode string `xml:"return_code,omitempty" json:"return_code"`
	ReturnMSG  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`

	ResultCode string `xml:"result_code,omitempty" json:"result_code"`     // 是	String(16)	SUCCESS/FAIL
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"` // 否	String(32)	SYSTEMERROR--系统错误
	ErrMSG     string `xml:"errmsg,omitempty" json:"errmsg,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

type ResponseMiniProgram struct {
	ResponseBase

	Msg     string `json:"msg,omitempty"` // 小程序直播的部分接口会把错误提示抛在msg字段
	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`
}

type ResponseOfficialAccount struct {
	ResponseBase

	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`
}

type ResponseOpenPlatform struct {
	ResponseBase

	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`
}

const (
	TYPE_RAW   = "raw"
	TYPE_MAP   = "map"
	TYPE_ARRAY = "array"
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
