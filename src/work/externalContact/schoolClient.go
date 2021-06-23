package externalContact

import "github.com/ArtisanCloud/go-wechat/src/kernel"

type SchoolClient struct {
	*kernel.BaseClient
}

func NewSchoolClient(app kernel.ApplicationInterface) *SchoolClient {
	return &SchoolClient{
		kernel.NewBaseClient(&app, nil),
	}
}