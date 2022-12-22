package transfer

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, *BatchClient, error) {

	client, err := NewClient(&app)
	if err != nil {
		return nil, nil, err
	}
	batchClient, err := NewBatchClient(&app)
	if err != nil {
		return nil, nil, err
	}

	return client, batchClient, nil

}
