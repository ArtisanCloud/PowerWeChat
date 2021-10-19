package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMeetingRoomGetBookingInfo struct {
	*response.ResponseWork

	BookingList []*power.HashMap `json:"booking_list"`
}
