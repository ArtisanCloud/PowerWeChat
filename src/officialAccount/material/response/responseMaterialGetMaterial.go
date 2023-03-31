package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type NewsItems struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int8   `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

type ResponseMaterialGetNews struct {
	response.ResponseOfficialAccount

	NewsItem []NewsItems `json:"news_item"`
}

// ------------------------------------

type ResponseMaterialGetVideo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownUrl     string `json:"down_url"`
}
