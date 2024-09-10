package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Link struct {
	LinkId     string `json:"link_id"`   // 链接ID
	LinkName   string `json:"link_name"` // 链接名称
	Url        string `json:"url"`
	SkipVerify bool   `json:"skip_verify"`
	CreateTime int64  `json:"create_time"`
}

type ResponseCreateLink struct {
	response.ResponseWork

	Link Link `json:"link"`
}
