package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	*response.ResponsePayment

	Status string `header:"status"`
}
