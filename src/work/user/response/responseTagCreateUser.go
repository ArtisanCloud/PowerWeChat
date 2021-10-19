package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseTagCreateUser struct {
	*response.ResponseWork

	InvalidUser  string `json:"invaliduser"`  // "userid1|userid2",
	InvalidParty string `json:"invalidparty"` // "partyid1|partyid2",
}
