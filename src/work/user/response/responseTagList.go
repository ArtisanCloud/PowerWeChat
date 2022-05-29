package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/user/request"
)

type ResponseTagList struct {
	*response.ResponseWork

	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
