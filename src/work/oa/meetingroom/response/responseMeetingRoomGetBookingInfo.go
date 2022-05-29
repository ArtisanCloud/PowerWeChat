package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMeetingRoomGetBookingInfo struct {
	*response.ResponseWork

	BookingList []*power.HashMap `json:"booking_list"`
}
