package sandbox

import (
	"github.com/ArtisanCloud/powerwechat/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) *Client {

	return NewClient(&app)

}
