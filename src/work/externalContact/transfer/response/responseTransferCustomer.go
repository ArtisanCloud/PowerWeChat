package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseTransferCustomer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
