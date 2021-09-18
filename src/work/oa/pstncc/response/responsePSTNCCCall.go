package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	*response.ResponseWork

	States []*power.HashMap `json:"states"`

}
