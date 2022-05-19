package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseOfficialAccount

	URL string `json:"url"`
}
