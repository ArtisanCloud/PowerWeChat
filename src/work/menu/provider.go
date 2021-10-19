package menu

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := NewClient(app)

	return client

}
