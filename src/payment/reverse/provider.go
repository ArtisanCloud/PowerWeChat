package reverse

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	return NewClient(&app)

}
