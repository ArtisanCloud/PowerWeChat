package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseShortGenKey struct {
	response.ResponseOfficialAccount

	ShortKey string `json:"short_key"`
}

type ResponseFetchShorten struct {
	response.ResponseOfficialAccount

	LongData      string `json:"long_data"`
	CreateTime    int    `json:"create_time"`
	ExpireSeconds int    `json:"expire_seconds"`
}
