package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/user/request"
)

type ResponseTagList struct {
	*response.ResponseWork

	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
