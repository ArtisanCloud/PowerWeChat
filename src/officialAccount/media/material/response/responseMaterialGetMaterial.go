package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMaterialGetMaterial struct {
	*response.ResponseOfficialAccount

	NewsItem []*power.HashMap `json:"news_item"`
}
