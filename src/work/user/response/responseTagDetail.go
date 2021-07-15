package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagDetail struct {
	response.ResponseWork
	TagName   string              `json:"tagname"`
	UserList  []*UserSimpleDetail `json:"userlist"`
	PartyList []int               `json:"partylist"`
}
