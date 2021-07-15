package base

import (
	"github.com/ArtisanCloud/power-wechat/src/payment"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

func RegisterProvider(app *payment.Payment) *Client {

	return &Client{
		kernel.NewBaseClient(app),
	}

}
