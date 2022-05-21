package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type RelationClient struct {
	*kernel.BaseClient
}


func NewRelationClient(app kernel.ApplicationInterface) *RelationClient {
	return &RelationClient{
		kernel.NewBaseClient(&app, nil),
	}
}