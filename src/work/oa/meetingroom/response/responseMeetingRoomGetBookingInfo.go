package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMeetingRoomGetBookingInfo struct {
	*response.ResponseWork

	BookingList []*power.HashMap `json:"booking_list"`
}
