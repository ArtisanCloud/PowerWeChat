package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	*response.ResponseWork

	ItemList []*power.HashMap `json:"item_list"`
}
