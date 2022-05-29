package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
)

type Transfer struct {
	*Message
}

func NewTransfer(account string) *Transfer {
	m := &Transfer{
		NewMessage(&power.HashMap{"account": account}),
	}

	m.Type = "transfer_customer_service"

	m.Properties = []string{
		"account",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Transfer) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		account := msg.GetString("account", "")
		if account != "" {
			return &object.HashMap{
				"TransInfo": &object.HashMap{
					"KfAccount": msg.GetString("account", ""),
				},
			}
		} else {
			return nil
		}

	}
}
