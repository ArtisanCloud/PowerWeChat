package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	*response.ResponseWork

	States []*power.HashMap `json:"states"`

}
