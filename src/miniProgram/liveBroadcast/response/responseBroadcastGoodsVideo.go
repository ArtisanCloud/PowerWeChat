package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	*response.ResponseMiniProgram

	URL int `json:"url"`
}
