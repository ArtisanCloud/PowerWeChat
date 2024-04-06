package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponsePublishSubmit struct {
	response.ResponseOfficialAccount

	PublishId uint64 `json:"publish_id"`
}

type ArticleItem struct {
	Idx        int    `json:"idx"`
	ArticleUrl string `json:"article_url"`
}

type ArticleDetail struct {
	Count int            `json:"count"`
	Item  []*ArticleItem `json:"item"`
}

type ResponsePublishGet struct {
	response.ResponseOfficialAccount

	PublishId     uint64         `json:"publish_id"`
	PublishStatus int            `json:"publish_status"`
	ArticleId     interface{}    `json:"article_id"`
	ArticleDetail *ArticleDetail `json:"article_detail"`
	FailIdx       []int          `json:"fail_idx"`
}

type ResponsePublishGetArticle struct {
	response.ResponseOfficialAccount

	NewsItem []*NewsItem `json:"news_item"`
}
