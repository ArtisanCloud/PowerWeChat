package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		jssdk.NewClient(&app),
	}

}
