package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	*response.ResponseMiniProgram

	URL int `json:"url"`
}
