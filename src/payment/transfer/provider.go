package transfer

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, *BatchClient) {

	client := NewClient(&app)
	batchClient := NewBatchClient(&app)

	return client, batchClient

}
