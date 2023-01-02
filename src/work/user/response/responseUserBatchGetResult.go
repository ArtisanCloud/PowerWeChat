package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserBatchGetResult struct {
	response.ResponseWork

	Status     int               `json:"status"`
	Type       string            `json:"type"`
	Total      int               `json:"total"`
	Percentage int               `json:"percentage"`
	Result     []*object.HashMap `json:"result"`
}
