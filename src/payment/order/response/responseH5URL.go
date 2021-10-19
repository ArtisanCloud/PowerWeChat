package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseH5URL struct {
	*response.ResponsePayment

	H5URL string `json:"h5_url"`
}
