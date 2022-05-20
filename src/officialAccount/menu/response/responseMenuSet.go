package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseOfficialAccount

}



type ResponseMenuCreateConditional struct {
	*response.ResponseOfficialAccount

	MenuID string `json:"menuid"`

}

