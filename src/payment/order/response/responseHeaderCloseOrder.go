package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	*response.ResponsePayment

	Status string `header:"status"`
}
