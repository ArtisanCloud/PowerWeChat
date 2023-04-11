package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type MsgMenu struct {
	*Message
}

func NewMsgMenu(items *power.HashMap) *MsgMenu {
	m := &MsgMenu{
		NewMessage(items),
	}

	m.Type = "msgmenu"

	m.Properties = []string{
		"head_content",
		"tail_content",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *MsgMenu) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		items := []*NewsItem{}
		getItem := msg.Get("items", nil)
		arrayItems := getItem.([]*NewsItem)
		for _, item := range arrayItems {
			items = append(items, item)
		}
		music := &object.HashMap{
			"MsgMenu": &object.HashMap{
				"Title":       msg.Get("head_content", ""),
				"Description": msg.Get("tail_content", ""),
			},
		}

		return music
	}
}
