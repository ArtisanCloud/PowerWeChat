package redpack

import (
	"github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) *Client {

	return NewClient(&app)

}
