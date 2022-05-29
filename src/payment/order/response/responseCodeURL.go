package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCodeURL struct {
	*response.ResponsePayment

	CodeURL string `json:"code_url"`
}
