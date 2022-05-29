package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	*response.ResponsePayment

	Status string `header:"status"`
}
