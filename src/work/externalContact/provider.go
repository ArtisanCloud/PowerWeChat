package externalContact

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *ContactWayClient) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := NewContactClient(app)

	return client, contactWayClient
}
