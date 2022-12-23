package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type DeviceEvent struct {
	*Message
}

func NewDeviceEvent(items *power.HashMap) *DeviceEvent {
	m := &DeviceEvent{
		NewMessage(items),
	}

	m.Type = "device_event"
	m.Properties = []string{
		"device_type",
		"device_id",
		"content",
		"session_id",
		"open_id",
	}

	return m
}
