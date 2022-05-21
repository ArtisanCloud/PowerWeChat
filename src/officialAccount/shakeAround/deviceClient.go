package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type DeviceClient struct {
	*kernel.BaseClient
}

func NewDeviceClient(app kernel.ApplicationInterface) *DeviceClient {
	return &DeviceClient{
		kernel.NewBaseClient(&app, nil),
	}
}
