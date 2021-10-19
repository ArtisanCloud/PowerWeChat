package jssdk

import "github.com/ArtisanCloud/powerwechat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return NewClient(&app)

}
