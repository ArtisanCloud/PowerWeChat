package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMaterialAddNews struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
