package jssdk

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/jssdk/response"
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

func (comp *Client) GetTicket(ctx context.Context) (*response2.ResponseGetTicket, error) {
	result := &response2.ResponseGetTicket{}

	params := &object.StringMap{
		"type": "agent_config",
	}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/ticket/get", params, nil, result)

	return result, err
}
