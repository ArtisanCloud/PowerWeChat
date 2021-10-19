package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseTagCreate struct {
	*response.ResponseWork

	TagID string `json:"tagid"`
}
