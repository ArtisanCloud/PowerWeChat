package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMaterialAddMaterial struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
	URL string `json:"url"`
}
