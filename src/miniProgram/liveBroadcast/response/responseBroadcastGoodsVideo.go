package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	*response.ResponseMiniProgram

	URL int `json:"url"`
}
