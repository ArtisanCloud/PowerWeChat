package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*power.HashMap `json:"button"`
}
