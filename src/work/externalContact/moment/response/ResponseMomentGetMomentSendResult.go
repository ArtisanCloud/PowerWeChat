package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMomentGetMomentSendResult struct {
	*response.ResponseWork

	CustomerList []*power.HashMap `json:"customer_list"`
	NextCursor   string           `json:"next_cursor"`
}
