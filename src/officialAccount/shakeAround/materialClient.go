package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type MaterialClient struct {
	*kernel.BaseClient
}

func NewMaterialClient(app kernel.ApplicationInterface) *MaterialClient {
	return &MaterialClient{
		kernel.NewBaseClient(&app, nil),
	}
}