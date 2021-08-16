package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*object.HashMap `json:"button"`
}

