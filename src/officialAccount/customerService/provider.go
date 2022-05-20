package customerService

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/customerService/session"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *session.Client) {

	client := NewClient(app)

	sessionClient := session.NewClient(app)

	return client, sessionClient
}
