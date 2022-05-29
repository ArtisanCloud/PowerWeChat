package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	*response.ResponseMiniProgram

	URL int `json:"url"`
}
