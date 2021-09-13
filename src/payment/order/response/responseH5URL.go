package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseH5URL struct {
	*response.ResponsePayment

	H5URL string `json:"h5_url"`
}
