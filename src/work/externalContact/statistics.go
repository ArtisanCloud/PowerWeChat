package externalContact

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type Statistics struct {
	*kernel.BaseClient
}

func NewStatisticsClient(app kernel.ApplicationInterface) *Statistics {
	return &Statistics{
		kernel.NewBaseClient(&app, nil),
	}
}
