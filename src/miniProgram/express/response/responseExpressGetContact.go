package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseExpressGetContact struct {
	*response2.ResponseMiniProgram

	WayBillID int              `json:"waybill_id"`
	Sender    []*power.HashMap `json:"sender"`
	Receiver  []*power.HashMap `json:"receiver"`
}
