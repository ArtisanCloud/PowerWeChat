package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
)

type Article struct {
	*Message


}

func NewArticle(items *power.HashMap) *Article {
	m := &Article{
		NewMessage(items),
	}

	m.Type = "mpnews"
	m.Properties = []string{
		"thumb_media_id",
		"author",
		"title",
		"content",
		"digest",
		"source_url",
		"show_cover",
	}

	m.JsonAliases = &object.HashMap{
		"content_source_url" : "source_url",
		"show_cover_pic" : "show_cover",
	}

	m.SetAttributes(&object.HashMap{
		"required":[]string{
			"thumb_media_id",
			"title",
			"content",
			"show_cover",
		},
	})

	return m
}

