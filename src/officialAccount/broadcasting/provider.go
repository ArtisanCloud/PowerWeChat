package broadcasting

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := &Client{
		kernel.NewBaseClient(&app, nil),
	}

	return client
}


