package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseApprovalGetData struct {
	*response.ResponseWork
	Count     int             `json:"count"`
	Total     int             `json:"total"`
	NextSPNum int             `json:"next_spnum"`
	Data      *object.HashMap `json:"data"`
}
