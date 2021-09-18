package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMaterialBatchGetMaterial struct {
	*response.ResponseOfficialAccount

	TotalCount int              `json:"total_count"`
	ItemCount  int              `json:"item_count"`
	Item       []*power.HashMap `json:"item"`
}
