package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	*response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
