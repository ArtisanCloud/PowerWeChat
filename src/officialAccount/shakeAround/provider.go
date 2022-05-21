package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),
		Device:     NewDeviceClient(app),
		Page:       NewPageClient(app),
		Material:   NewMaterialClient(app),
		Group:      NewGroupClient(app),
		Relation:   NewRelationClient(app),
		Stats:      NewStatsClient(app),
	}

	return client
}
