package menu

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := NewClient(app)

	return client

}
