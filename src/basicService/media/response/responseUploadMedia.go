package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseUploadMedia struct {
	response.ResponseOfficialAccount

	Item         []*power.HashMap `json:"item"`
	Type         string           `json:"type"`
	MediaID      string           `json:"media_id"`
	ThumbMediaID string           `json:"thumb_media_id,omitempty"`
	CreatedAt    int              `json:"created_at"`
}
