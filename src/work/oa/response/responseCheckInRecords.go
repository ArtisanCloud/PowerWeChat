package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckInRecords struct {
	*response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
