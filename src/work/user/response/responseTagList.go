package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/request"
)

type ResponseTagList struct {
	*response.ResponseWork

	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
