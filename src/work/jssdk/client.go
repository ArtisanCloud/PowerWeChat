package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
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

	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	client.TicketEndpoint = baseURI + "/cgi-bin/get_jsapi_ticket"

	return client, nil
}

func (comp *Client) GetAppID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("corp_id", "")
}

func (comp *Client) GetAgentConfigArray() {

}
