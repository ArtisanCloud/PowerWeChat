package sandbox

import (
	kernel2 "github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"

)

func RegisterProvider(app *payment.Payment) (*Client) {

	client := &Client{
		kernel.NewBaseClient(app),
		&kernel2.InteractsWithCache{},
	}

	return client
}
