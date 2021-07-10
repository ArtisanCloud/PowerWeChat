package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/user/request"
)

type ResponseTagList struct {
	response.ResponseWX
	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
