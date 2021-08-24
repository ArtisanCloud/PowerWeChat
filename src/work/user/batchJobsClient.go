package user

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type BatchJobsClient struct {
	*kernel.BaseClient
}

func NewBatchJobsClient(app kernel.ApplicationInterface) *BatchJobsClient {
	return &BatchJobsClient{
		kernel.NewBaseClient(&app, nil),
	}
}
