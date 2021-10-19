package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseTransferResult struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
	NextCursor string `json:"next_cursor"`
}
