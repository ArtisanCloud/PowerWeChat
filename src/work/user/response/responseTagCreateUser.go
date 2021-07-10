package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagCreateUser struct {
	response.ResponseWX
	InvalidUser  string `json:"invaliduser"`  // "userid1|userid2",
	InvalidParty string `json:"invalidparty"` // "partyid1|partyid2",
}
