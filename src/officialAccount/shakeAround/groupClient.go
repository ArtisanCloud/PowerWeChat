package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type GroupClient struct {
	*kernel.BaseClient
}


func NewGroupClient(app kernel.ApplicationInterface) *GroupClient {
	return &GroupClient{
		kernel.NewBaseClient(&app, nil),
	}
}