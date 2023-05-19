package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	jssdkClient, err := jssdk.NewClient(app)
	if err != nil {
		return nil, err
	}
	client := &Client{
		jssdkClient,
	}

	client.TicketEndpoint = "cgi-bin/get_jsapi_ticket"

	return client, nil
}


func (comp *Client) GetAgentConfigArray() {

}
