package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMomentGetMomentSendResult struct {
	*response.ResponseWork

	CustomerList []*power.HashMap `json:"customer_list"`
	NextCursor   string           `json:"next_cursor"`
}
