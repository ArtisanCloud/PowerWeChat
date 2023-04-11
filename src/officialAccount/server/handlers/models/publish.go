package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

// https://developers.weixin.qq.com/doc/offiaccount/Publish/Callback_on_finish.html

const (
	CALLBACK_EVENT_PUBLISH = "PUBLISHJOBFINISH"
)

type EventPublish struct {
	contract.EventInterface
	models.CallbackMessageHeader

	PublishEventInfo PublishEventInfo `json:"PublishEventInfo"`
}

type PublishEventInfo struct {
	PublishID     string        `json:"publish_id"`
	PublishStatus string        `json:"publish_status"`
	ArticleID     string        `json:"article_id"`
	ArticleDetail ArticleDetail `json:"article_detail"`
}

type ArticleDetail struct {
	Count string `json:"count"`
	Item  Item   `json:"item"`
}

type Item struct {
	Idx        string `json:"idx"`
	ArticleURL string `json:"article_url"`
}
