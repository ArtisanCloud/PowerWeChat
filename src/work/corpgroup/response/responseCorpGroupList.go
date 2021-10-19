package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCorpGroupListAPPShareInfo struct {
	*response.ResponseWork
	CorpList []*power.HashMap `json:"corp_list"`
}
