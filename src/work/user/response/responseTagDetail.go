package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseTagDetail struct {
	*response.ResponseWork

	TagName   string              `json:"tagname"`
	UserList  []*UserSimpleDetail `json:"userlist"`
	PartyList []int               `json:"partylist"`
}
