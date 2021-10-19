package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCheckInRecords struct {
	*response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
