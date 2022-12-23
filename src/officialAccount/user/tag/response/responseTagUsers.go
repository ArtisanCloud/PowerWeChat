package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseTagUsers struct {
	*response.ResponseOfficialAccount

	OpenIDList []string `json:"openid_list"`
	TagID      int      `json:"tagid"`
}

type ResponseUntagUsers struct {
	*response.ResponseOfficialAccount

	OpenIDList []string `json:"openid_list"`
	TagID      int      `json:"tagid"`
}
