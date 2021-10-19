package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCodeURL struct {
	*response.ResponsePayment

	CodeURL string `json:"code_url"`
}
