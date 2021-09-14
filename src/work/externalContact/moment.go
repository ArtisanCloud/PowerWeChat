package externalContact

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type Moment struct {
	*kernel.BaseClient
}

func NewMomentClient(app kernel.ApplicationInterface) *Moment {
	return &Moment{
		kernel.NewBaseClient(&app, nil),
	}
}
