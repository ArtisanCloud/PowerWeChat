package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	*response.ResponseWork

	URL string `json:"url"`
}
