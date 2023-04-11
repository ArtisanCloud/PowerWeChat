package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	response.ResponsePayment

	Status string `header:"status"`
}
