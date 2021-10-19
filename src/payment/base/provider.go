package base

import (
	"github.com/ArtisanCloud/powerwechat/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) *Client {

	return &Client{
		kernel.NewBaseClient(&app),
	}

}
