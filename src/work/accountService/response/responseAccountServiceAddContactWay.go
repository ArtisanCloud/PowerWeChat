package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	*response.ResponseWork

	URL string `json:"url"`
}
