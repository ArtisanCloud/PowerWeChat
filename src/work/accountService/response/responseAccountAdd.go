package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAccountAdd struct {
	*response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
