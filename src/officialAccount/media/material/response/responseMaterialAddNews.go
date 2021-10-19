package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMaterialAddNews struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
