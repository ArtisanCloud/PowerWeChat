package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckInDatas struct {
	*response.ResponseWork

	Datas []*object.HashMap `json:"datas"`
}
