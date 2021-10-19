package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseTagCreate struct {
	*response.ResponseWork

	TagID string `json:"tagid"`
}
