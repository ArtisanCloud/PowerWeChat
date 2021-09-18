package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMaterialGetMaterial struct {
	*response.ResponseOfficialAccount

	NewsItem []*power.HashMap `json:"news_item"`
}
