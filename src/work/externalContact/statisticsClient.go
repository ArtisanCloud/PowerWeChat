package externalContact

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type StatisticsClient struct {
	*kernel.BaseClient
}

func NewStatisticsClient(app kernel.ApplicationInterface) *StatisticsClient {
	return &StatisticsClient{
		kernel.NewBaseClient(&app, nil),
	}
}