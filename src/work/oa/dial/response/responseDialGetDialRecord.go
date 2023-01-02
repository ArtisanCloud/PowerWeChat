package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
