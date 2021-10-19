package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseWork

	URL string `json:"url"`
}
