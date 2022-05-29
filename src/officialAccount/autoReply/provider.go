package autoReply

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := &Client{
		kernel.NewBaseClient(&app, nil),
	}

	return client
}
