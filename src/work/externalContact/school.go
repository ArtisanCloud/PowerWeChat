package externalContact

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type School struct {
	*kernel.BaseClient
}

func NewSchoolClient(app kernel.ApplicationInterface) *School {
	return &School{
		kernel.NewBaseClient(&app, nil),
	}
}
