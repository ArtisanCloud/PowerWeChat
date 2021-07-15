package base

import (
	kernel2 "github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

func RegisterProvider(app *kernel2.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(app),
	}

}
