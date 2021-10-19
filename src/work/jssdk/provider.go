package jssdk

import (
	"github.com/ArtisanCloud/powerwechat/src/basicService/jssdk"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		jssdk.NewClient(&app),
	}

}
