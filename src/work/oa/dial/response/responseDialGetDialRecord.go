package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	*response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
