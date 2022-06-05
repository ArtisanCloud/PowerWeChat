package card

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),
	}

	return client
}
