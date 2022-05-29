package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*power.HashMap `json:"button"`
}
