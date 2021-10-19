package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUploadMedia struct {
	*response.ResponseOfficialAccount

	Item      []*power.HashMap `json:"item"`
	Type      string           `json:"type"`
	MediaID   string           `json:"media_id"`
	CreatedAt int              `json:"created_at"`
}
