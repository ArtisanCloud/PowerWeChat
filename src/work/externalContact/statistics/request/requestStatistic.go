package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestStatistic struct {
	DayBeginTime int64          `json:"day_begin_time" `
	DayEndTime   int64          `json:"day_end_time"`
	OwnerFilter  *power.HashMap `json:"owner_filter"`
	OrderBy      int            `json:"order_by"`
	OrderAsc     int            `json:"order_asc"`
	Offset       int            `json:"offset"`
	Limit        int            `json:"limit"`
}
