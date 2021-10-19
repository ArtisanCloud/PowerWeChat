package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseWork

	URL string `json:"url"`
}
