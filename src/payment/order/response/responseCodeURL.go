package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCodeURL struct {
	*response.ResponsePayment

	CodeURL string `json:"code_url"`
}
