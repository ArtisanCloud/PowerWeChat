package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCheckInRecords struct {
	*response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
