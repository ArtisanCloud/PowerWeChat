package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckInDatas struct {
	*response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
