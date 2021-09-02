package externalContact

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type MomentClient struct {
	*kernel.BaseClient
}

func NewMomentClient(app kernel.ApplicationInterface) *MomentClient {
	return &MomentClient{
		kernel.NewBaseClient(&app, nil),
	}
}
