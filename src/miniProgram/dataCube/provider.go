package dataCube

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(&app, nil),
	}

}
