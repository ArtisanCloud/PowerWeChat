package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type News struct {
	*Message
}

func NewNews(items []*object.HashMap) *News {
	m := &News{
		NewMessage(&power.HashMap{"items": items}),
	}
	m.Type = "news"
	m.OverrideToXmlArray()
	m.OverridePropertiesToArray()

	return m
}

func (msg *News) OverridePropertiesToArray() {

	msg.PropertiesToArray = func(data *object.HashMap, aliases *object.HashMap) (*object.HashMap, error) {
		arrayItems := msg.Get("items", nil).([]*object.HashMap)
		title := msg.Get("title", nil).(string)
		description := msg.Get("description", nil).(string)
		picUrl := msg.Get("picUrl", nil).(string)
		url := msg.Get("url", nil).(string)
		arrayMapItems := []*object.HashMap{}
		for _, item := range arrayItems {
			arrayMapItems = append(arrayMapItems, item)
		}

		return &object.HashMap{
			"articles":    arrayMapItems,
			"title":       title,
			"description": description,
			"picUrl":      picUrl,
			"url":         url,
		}, nil
	}

}

// Override ToXmlArray
func (msg *News) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		items := []*object.HashMap{}

		getItem := msg.Get("items", nil)
		title := msg.Get("title", nil).(string)
		description := msg.Get("description", nil).(string)
		picUrl := msg.Get("picUrl", nil).(string)
		url := msg.Get("url", nil).(string)
		if getItem != nil {
			arrayItems := getItem.([]*object.HashMap)
			for _, item := range arrayItems {
				//newItems := NewNewsItem(item)
				items = append(items, &object.HashMap{
					"item": item,
				})
			}
		}

		return &object.HashMap{
			"ArticleCount": len(items),
			"Articles":     items,
			"Title":        title,
			"Description":  description,
			"PicUrl":       picUrl,
			"Url":          url,
		}
	}
}
