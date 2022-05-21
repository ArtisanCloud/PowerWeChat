package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type StatsClient struct {
	*kernel.BaseClient
}


func NewStatsClient(app kernel.ApplicationInterface) *StatsClient {
	return &StatsClient{
		kernel.NewBaseClient(&app, nil),
	}
}