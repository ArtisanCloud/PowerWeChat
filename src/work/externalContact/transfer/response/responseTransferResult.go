package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseTransferResult struct {
	*response.ResponseWork

	Customer   []*power.HashMap `json:"customer"`
	NextCursor string           `json:"next_cursor"`
}
