package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLinkCorpGetUserList struct {
	*response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`

}
