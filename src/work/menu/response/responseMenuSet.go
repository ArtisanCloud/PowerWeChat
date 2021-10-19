package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*power.HashMap `json:"button"`
}
