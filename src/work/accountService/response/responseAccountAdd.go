package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAccountAdd struct {
	*response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
