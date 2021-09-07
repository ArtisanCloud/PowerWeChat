package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCustomerServiceMessageGetTempMedia struct {
	*response.ResponseMiniProgram
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}
