package user

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*BatchJobsClient,
	*LinkedCorpClient,
	*TagClient,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	UserBatchJobsClient := NewBatchJobsClient(app)
	UserLinkedCorpClient := NewLinkedCorpClient(app)
	UserTagClient := NewTagClient(app)

	return client,
		UserBatchJobsClient,
		UserLinkedCorpClient,
		UserTagClient
}
