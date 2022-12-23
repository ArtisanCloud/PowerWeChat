package miniProgram

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

type Client struct {
	*kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
	}, nil
}
