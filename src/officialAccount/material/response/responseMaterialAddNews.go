package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMaterialAddNews struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
