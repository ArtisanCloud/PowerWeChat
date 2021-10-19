package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUserBatchGetResult struct {
	*response.ResponseWork

	Status int `json:"status"`
	Type string `json:"type"`
	Total int `json:"total"`
	Percentage int `json:"percentage"`
	Result []*object.HashMap `json:"result"`

}
