package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type ShakeAroundClient struct {
	*kernel.BaseClient
}


func NewShakeAroundClient(app kernel.ApplicationInterface) *DeviceClient {
	return &DeviceClient{
		kernel.NewBaseClient(&app, nil),
	}
}