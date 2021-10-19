package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAccountAdd struct {
	*response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
