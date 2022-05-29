package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseWork

	URL string `json:"url"`
}
