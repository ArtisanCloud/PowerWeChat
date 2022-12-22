package base

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	baseClient, err := kernel.NewBaseClient(&app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil

}
