package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*ContactWay,
	*Statistics,
	*Message,
	*School,
	*Moment,
	*MessageTemplate,
	*GroupChat,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := NewContactClient(app)

	statistics := NewStatisticsClient(app)
	message := NewMessageClient(app)
	school := NewSchoolClient(app)
	moment := NewMomentClient(app)
	messageTemplate := NewMessageTemplate(app)
	groupChat := NewGroupChat(app)

	return client,
		contactWayClient,
		statistics,
		message,
		school,
		moment,
		messageTemplate,
		groupChat
}
