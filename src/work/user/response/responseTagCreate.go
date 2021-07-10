package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWX
	TagID string `json:"tagid"`
}
