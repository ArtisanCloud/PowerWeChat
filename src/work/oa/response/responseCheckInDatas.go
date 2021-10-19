package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCheckInDatas struct {
	*response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
