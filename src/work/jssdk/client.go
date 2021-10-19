package jssdk

import (
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app *kernel.ApplicationInterface) *Client {
	client := &Client{
		jssdk.NewClient(app),
	}

	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	client.TicketEndpoint = baseURI + "/cgi-bin/get_jsapi_ticket"

	return client
}



func (comp *Client) GetAppID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("corp_id", "")
}

func (comp *Client)GetAgentConfigArray()  {
	
}