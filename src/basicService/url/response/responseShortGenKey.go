package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseShortGenKey struct {
	*response.ResponseOfficialAccount

	ShortKey string `json:"short_key"`
}