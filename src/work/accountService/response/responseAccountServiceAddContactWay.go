package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	*response.ResponseWork

	URL string `json:"url"`
}
