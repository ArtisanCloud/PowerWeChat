package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMaterialGetMaterialCount struct {
	*response.ResponseOfficialAccount

	VoiceCount string `json:"voice_count"`
	VideoCount string `json:"video_count"`
	ImageCount string `json:"image_count"`
	NewsCount  string `json:"news_count"`
	MediaID    string `json:"media_id"`
}
