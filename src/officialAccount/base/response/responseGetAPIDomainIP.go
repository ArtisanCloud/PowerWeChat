package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	*response.ResponseOfficialAccount

	DNS  []power.StringMap `json:"dns"`
	Ping []power.StringMap `json:"ping"`
}
