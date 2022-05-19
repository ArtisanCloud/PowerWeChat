package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseUserTags struct {
	*response.ResponseOfficialAccount
	TagIDList []int `json:"tagid_list"`
}

// ------------

type Data struct {
	Openid []string `json:"openid"`
}

type ResponseUserOfTag struct {
	*response.ResponseOfficialAccount

	Count      int    `json:"count"`
	Data       *Data  `json:"data"`
	NextOpenid string `json:"next_openid"`
}
