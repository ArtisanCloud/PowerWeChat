package menu

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := NewClient(app)

	return client

}
