package card

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	return client, nil
}
