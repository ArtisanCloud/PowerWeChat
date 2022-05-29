package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseLinkCorpMessageSend struct {
	*response.ResponseWork

	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}
