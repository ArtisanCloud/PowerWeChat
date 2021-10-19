package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMaterialBatchGetMaterial struct {
	*response.ResponseOfficialAccount

	TotalCount int              `json:"total_count"`
	ItemCount  int              `json:"item_count"`
	Item       []*power.HashMap `json:"item"`
}
