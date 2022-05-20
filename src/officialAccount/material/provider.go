package material

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := NewClient(app)
	return client
}
