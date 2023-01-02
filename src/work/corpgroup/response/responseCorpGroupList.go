package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpGroupListAPPShareInfo struct {
	response.ResponseWork
	CorpList []*power.HashMap `json:"corp_list"`
}
