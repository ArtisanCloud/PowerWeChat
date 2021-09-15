package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseServicerList struct {
	*response.ResponseWork

	ServicerList []*power.HashMap `json:"servicer_list"`
}
