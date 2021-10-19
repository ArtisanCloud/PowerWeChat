package response

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUserExportGetResult struct {
	*response.ResponseWork

	Status int `json:"status"`
	DataList []*object.HashMap `json:"data_list"`

}
