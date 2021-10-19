package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCheckInDatas struct {
	*response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
