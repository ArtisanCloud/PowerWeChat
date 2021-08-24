package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	*response.ResponseWork

	ItemList []*object.HashMap `json:"item_list"`
}

