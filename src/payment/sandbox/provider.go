package sandbox

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"

)

func RegisterProvider(app kernel.ApplicationInterface) (*Client) {

	client := &Client{
		kernel.NewBaseClient(&app, nil),
	}

	return client
}
