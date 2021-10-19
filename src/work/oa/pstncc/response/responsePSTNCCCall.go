package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	*response.ResponseWork

	States []*power.HashMap `json:"states"`

}
