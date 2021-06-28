package response

import (
	"net/http"
)

type ResponseWX struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}


func (res *ResponseWX) GetBody() *http.ResponseWriter {
	return nil
}
func (res *ResponseWX) GetHeaders() *http.ResponseWriter {
	return nil
}

func (res *ResponseWX) GetStatusCode() int {
	return 200
}
