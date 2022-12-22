package store

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		baseClient,
	}

	return client, nil
}
