package messages

import (
	"encoding/base64"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type DeviceText struct {
	*Message
}

func NewDeviceText(items *power.HashMap) *DeviceText {
	m := &DeviceText{
		NewMessage(items),
	}

	m.Type = "device_text"

	m.Properties = []string{
		"device_type",
		"device_id",
		"content",
		"session_id",
		"open_id",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *DeviceText) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		content := base64.StdEncoding.EncodeToString([]byte(msg.GetString("content", "")))

		return &object.HashMap{
			"DeviceType": msg.Get("device_type", nil),
			"DeviceID":   msg.Get("device_id", nil),
			"SessionID":  msg.Get("session_id", nil),
			"Content":    content,
		}
	}
}
