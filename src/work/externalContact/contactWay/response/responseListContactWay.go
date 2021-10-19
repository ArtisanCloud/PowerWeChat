package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseListContactWay struct {
	*response.ResponseWork

	ContactWay []*power.HashMap `json:"contact_way"`
}
