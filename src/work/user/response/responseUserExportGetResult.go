package response

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseUserExportGetResult struct {
	*response.ResponseWork

	Status   int               `json:"status"`
	DataList []*object.HashMap `json:"data_list"`
}
