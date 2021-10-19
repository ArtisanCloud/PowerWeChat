package request

import "github.com/ArtisanCloud/powerwechat/src/kernel/power"

type RequestMeetingRoomEdit struct {
	MeetingRoomID int              `json:"meetingroom_id"`
	Name          string           `json:"name"`
	Capacity      int              `json:"capacity"`
	City          string           `json:"city"`
	Building      string           `json:"building"`
	Floor         string           `json:"floor"`
	Equipment     []int            `json:"equipment"`
	Coordinate    *power.StringMap `json:"coordinate"`
}
