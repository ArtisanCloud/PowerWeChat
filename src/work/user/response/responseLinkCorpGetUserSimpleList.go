package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	response.ResponseWork
	UserList []*power.HashMap `json:"userlist"`

}
