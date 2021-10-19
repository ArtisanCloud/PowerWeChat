package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCodeURL struct {
	*response.ResponsePayment

	CodeURL string `json:"code_url"`
}
