package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseExternalPayGetBillList struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	BillList   []*power.HashMap `json:"bill_list"`
}
