package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseBlacklist struct {
	*response.ResponseOfficialAccount

	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}
