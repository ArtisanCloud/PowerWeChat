package externalContact

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*ContactWayClient,
	*StatisticsClient,
	*MessageClient,
	*SchoolClient,
	*MomentClient,
	*MessageTemplateClient,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := NewContactClient(app)

	ExternalContactStatistics := NewStatisticsClient(app)
	ExternalContactMessage := NewMessageClient(app)
	School := NewSchoolClient(app)
	ExternalContactMoment := NewMomentClient(app)
	ExternalContactMessageTemplate := NewMessageTemplateClient(app)

	return client,
		contactWayClient,
		ExternalContactStatistics,
		ExternalContactMessage,
		School,
		ExternalContactMoment,
		ExternalContactMessageTemplate
}
