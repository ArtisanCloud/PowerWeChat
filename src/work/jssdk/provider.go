package jssdk

import (
	"github.com/ArtisanCloud/power-wechat/src/basicService/jssdk"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		jssdk.NewClient(&app),
	}

}
