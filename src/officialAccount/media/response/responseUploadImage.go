package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseOfficialAccount

	URL string `json:"url"`
}
