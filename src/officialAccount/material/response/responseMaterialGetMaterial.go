package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMaterialGetMaterial struct {
	*response.ResponseOfficialAccount

	NewsItem []*power.HashMap `json:"news_item"`
}
