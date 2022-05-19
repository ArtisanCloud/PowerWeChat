package messages

import "github.com/ArtisanCloud/PowerLibs/object"

type Article struct {
	*Message


}

func NewArticle(items *object.HashMap) *Article {
	m := &Article{
		NewMessage(&object.HashMap{"items": items}),
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



	return m
}

