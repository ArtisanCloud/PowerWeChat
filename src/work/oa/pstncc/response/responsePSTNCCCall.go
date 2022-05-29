package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	*response.ResponseWork

	States []*power.HashMap `json:"states"`
}
