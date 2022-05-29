package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	*response.ResponseWork

	URL string `json:"url"`
}
