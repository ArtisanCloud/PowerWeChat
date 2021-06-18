package message

import "github.com/ArtisanCloud/go-wechat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(&app, nil),
	}

}
