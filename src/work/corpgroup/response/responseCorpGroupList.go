package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpGroupListAPPShareInfo struct {
	*response.ResponseWork
	CorpList []*power.HashMap `json:"corp_list"`
}
