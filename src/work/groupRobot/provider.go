package groupRobot

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Messager) {

	client := NewClient(app)

	messager := NewMessager(client)

	return client, messager

}
