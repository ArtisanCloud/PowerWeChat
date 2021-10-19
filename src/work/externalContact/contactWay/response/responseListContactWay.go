package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseListContactWay struct {
	*response.ResponseWork

	ContactWay []*power.HashMap `json:"contact_way"`
}
