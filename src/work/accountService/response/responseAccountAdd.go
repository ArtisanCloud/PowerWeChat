package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAccountAdd struct {
	*response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
