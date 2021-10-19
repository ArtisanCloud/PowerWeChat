package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	*response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
