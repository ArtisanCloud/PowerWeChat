package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	*response.ResponseWork

	Strategy   *power.HashMap `json:"strategy"`

}