package immediateDelivery

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(&app, nil),
	}

}
