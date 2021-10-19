package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLinkCorpGetUserList struct {
	*response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`

}
