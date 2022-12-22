package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type News struct {
	*Message
}

func NewNews(items *object.HashMap) *News {
	m := &News{
		NewMessage(&power.HashMap{"items": items}),
	}
	m.OverrideToXmlArray()

	return m
}

func (msg *News) PropertiesToArray(data power.HashMap, aliases power.HashMap) *power.HashMap {

	arrayItems := msg.Get("items", nil).([]*NewsItem)
	arrayMapItems := []*object.HashMap{}
	for _, item := range arrayItems {
		arrayMapItems = append(arrayMapItems, item.ToJsonArray())
	}

	return &power.HashMap{
		"articles": arrayMapItems,
	}
}

// Override ToXmlArray
func (msg *News) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		items := []*NewsItem{}

		getItem := msg.Get("items", nil)
		if getItem != nil {
			arrayItems := getItem.([]*NewsItem)
			for _, item := range arrayItems {
				items = append(items, item)
			}
		}

		return &object.HashMap{
			"ArticleCount": len(items),
			"Articles":     items,
		}
	}
}
