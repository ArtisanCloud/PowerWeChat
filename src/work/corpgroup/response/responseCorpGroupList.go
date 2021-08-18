package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpGroupListAPPShareInfo struct {
	*response.ResponseWork
	CorpList []*object.HashMap `json:"corp_list"`
}

