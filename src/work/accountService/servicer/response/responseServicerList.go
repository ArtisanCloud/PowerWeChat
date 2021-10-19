package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseServicerList struct {
	*response.ResponseWork

	ServicerList []*power.HashMap `json:"servicer_list"`
}
