package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*ContactWay,
	*StatisticsClient,
	*Message,
	*SchoolClient,
	*Moment,
	*MessageTemplate,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := NewContactClient(app)

	ExternalContactStatistics := NewStatisticsClient(app)
	ExternalContactMessage := NewMessageClient(app)
	School := NewSchoolClient(app)
	ExternalContactMoment := NewMomentClient(app)
	ExternalContactMessageTemplate := NewMessageTemplate(app)

	return client,
		contactWayClient,
		ExternalContactStatistics,
		ExternalContactMessage,
		School,
		ExternalContactMoment,
		ExternalContactMessageTemplate
}
