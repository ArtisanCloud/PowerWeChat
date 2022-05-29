package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	*response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
