package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseServicerList struct {
	response.ResponseWork

	ServicerList []*power.HashMap `json:"servicer_list"`
}
