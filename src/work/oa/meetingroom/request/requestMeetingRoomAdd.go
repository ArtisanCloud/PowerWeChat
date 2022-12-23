package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestMeetingRoomAdd struct {
	Name       string           `json:"name"`
	Capacity   int              `json:"capacity"`
	City       string           `json:"city"`
	Building   string           `json:"building"`
	Floor      string           `json:"floor"`
	Equipment  []int            `json:"equipment"`
	Coordinate *power.StringMap `json:"coordinate"`
}
