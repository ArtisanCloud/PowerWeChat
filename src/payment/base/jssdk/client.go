package jssdk

import (
	"github.com/ArtisanCloud/power-wechat/src/basicService/jssdk"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

type Client struct {
	*jssdk.Client


}

func NewClient(app *kernel.ApplicationPaymentInterface) *Client {
	return &Client{
		jssdk.NewClient(app),
	}
}

func (comp *Client) GetKey() (string, error) {
	return "",nil
}
