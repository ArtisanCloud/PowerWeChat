package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		jssdk.NewClient(&app),
	}

}
