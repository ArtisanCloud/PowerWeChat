package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*ShakeAround, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, err
	}
	deviceClientClient, err := NewDeviceClient(app)
	if err != nil {
		return nil, err
	}
	pageClientClient, err := NewPageClient(app)
	if err != nil {
		return nil, err
	}
	materialClientClient, err := NewMaterialClient(app)
	if err != nil {
		return nil, err
	}
	groupClientClient, err := NewGroupClient(app)
	if err != nil {
		return nil, err
	}
	relationClientClient, err := NewRelationClient(app)
	if err != nil {
		return nil, err
	}
	statsClientClient, err := NewStatsClient(app)
	if err != nil {
		return nil, err
	}

	shakeAround := &ShakeAround{
		Client:   client,
		Device:   deviceClientClient,
		Page:     pageClientClient,
		Material: materialClientClient,
		Group:    groupClientClient,
		Relation: relationClientClient,
		Stats:    statsClientClient,
	}

	return shakeAround, nil
}
