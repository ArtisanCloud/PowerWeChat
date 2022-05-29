package contentSecurity

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

type Client struct {
	*kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) *Client {

	client := &Client{
		BaseClient: kernel.NewBaseClient(app, nil),
	}

	client.BaseClient.HttpRequest.BaseURI = "https://api.weixin.qq.com/wxa/"

	return client
}
