package jssdk

import "github.com/ArtisanCloud/power-wechat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return NewClient(&app)

}
