package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	jssdkClient, err := jssdk.NewClient(&app)
	if err != nil {
		return nil, err
	}
	return &Client{
		jssdkClient,
	}, nil

}
