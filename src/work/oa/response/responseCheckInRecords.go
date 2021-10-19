package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCheckInRecords struct {
	*response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
