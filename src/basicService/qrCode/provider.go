package qrCode

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return NewClient(&app)

}
