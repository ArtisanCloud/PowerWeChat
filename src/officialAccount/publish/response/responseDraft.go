package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseDraftAdd struct {
	response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}

type NewsItem struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
	Url                string `json:"url"`
	IsDeleted          bool   `json:"is_deleted"`
}

type ResponseDraftGet struct {
	response.ResponseOfficialAccount

	NewsItem []*NewsItem `json:"news_item"`
}

type ResponseDraftCount struct {
	TotalCount int `json:"total_count"`
}

type Content struct {
	NewsItem []*NewsItem `json:"news_item"`
}

type Item struct {
	MediaId    string   `json:"media_id"`
	ArticleId  string   `json:"article_id"`
	Content    *Content `json:"content"`
	UpdateTime int64    `json:"update_time"`
}

type ResponseBatchGet struct {
	response.ResponseOfficialAccount

	TotalCount int     `json:"total_count"`
	ItemCount  int     `json:"item_count"`
	Item       []*Item `json:"item"`
}

type ResponseCheckSwitch struct {
	response.ResponseOfficialAccount

	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	IsOpen     int `json:"is_open"`
}
