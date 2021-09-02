package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	*response.ResponseWork

	ItemList []*power.HashMap `json:"item_list"`
}

