package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseUploadImage struct {
	response.ResponseWork

	URL string `json:"url"`
}
