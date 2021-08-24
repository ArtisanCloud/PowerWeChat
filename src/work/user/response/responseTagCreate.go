package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWork
	TagID string `json:"tagid"`
}
