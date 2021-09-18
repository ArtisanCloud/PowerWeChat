package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	*response.ResponseOfficialAccount
	DNS []power.StringMap `json:"dns"`
	Ping []power.StringMap `json:"ping"`
}
