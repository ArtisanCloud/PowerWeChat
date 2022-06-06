package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/card/request"
)

type ResponseCardCreate struct {
	*response.ResponseOfficialAccount

	CardID string `json:"card_id"`
}

type ResponseCardGet struct {
	*response.ResponseOfficialAccount

	Card request.Card `json:"card"`
}
