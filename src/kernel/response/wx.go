package response

import "github.com/ArtisanCloud/go-libs/object"

type ResponseWX struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

func (res *ResponseWX)GetBody() *object.HashMap{
	return nil
}
func (res *ResponseWX)GetHeaders() *object.HashMap{
	return nil
}

func (res *ResponseWX)GetStatusCode() int{
	return 200
}