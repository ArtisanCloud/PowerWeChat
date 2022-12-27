package contentSecurity

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

type Client struct {
	BaseClient *kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {

	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	client.BaseClient.BaseURI = "https://api.weixin.qq.com/wxa/"

	return client, nil
}
