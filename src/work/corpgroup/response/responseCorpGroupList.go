package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCorpGroupListAPPShareInfo struct {
	*response.ResponseWork
	CorpList []*power.HashMap `json:"corp_list"`
}
