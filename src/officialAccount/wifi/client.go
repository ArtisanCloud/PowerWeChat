package wifi

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}