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

	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	client.TicketEndpoint = baseURI + "/cgi-bin/get_jsapi_ticket"

	client.OverrideGetAppID()

	return client, nil
}

func (comp *Client) OverrideGetAppID() {
	comp.GetAppID = func() string {
		config := (*comp.BaseClient.App).GetConfig()
		return config.GetString("corp_id", "")
	}
}

func (comp *Client) GetAgentConfigArray() {

}
