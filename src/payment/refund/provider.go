package refund

import (
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) *Client {

	return NewClient(&app)

}
