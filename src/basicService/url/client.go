package url

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type Client struct {
	*kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) *Client {

	return &Client{
		BaseClient: kernel.NewBaseClient(app, nil),

		BaseUri: "https://api.weixin.qq.com/",
	}

}

