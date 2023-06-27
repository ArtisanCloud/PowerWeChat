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
		arrayMapItems := []*object.HashMap{}
		for _, item := range arrayItems {
			arrayMapItems = append(arrayMapItems, item)
		}

		return &object.HashMap{
			"articles": arrayMapItems,
		}, nil
	}

}

// Override ToXmlArray
func (msg *News) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		items := []*object.HashMap{}

		getItem := msg.Get("items", nil)
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
		}
	}
}
