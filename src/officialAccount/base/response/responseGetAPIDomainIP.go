package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	*response.ResponseOfficialAccount
	DNS []power.StringMap `json:"dns"`
	Ping []power.StringMap `json:"ping"`
}
