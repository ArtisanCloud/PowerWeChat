package jssdk

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)

}
