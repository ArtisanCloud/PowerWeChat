package jssdk

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return NewClient(&app)

}
