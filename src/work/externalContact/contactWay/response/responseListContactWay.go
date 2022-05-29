package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseListContactWay struct {
	*response.ResponseWork

	ContactWay []*power.HashMap `json:"contact_way"`
}
