package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/message"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/moment"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/school"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/statistics"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*contactWay.ContactWay,
	*statistics.Statistics,
	*message.Message,
	*school.School,
	*moment.Moment,
	*messageTemplate.MessageTemplate,
	*groupChat.GroupChat,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := contactWay.NewContactClient(app)

	statistics := statistics.NewStatisticsClient(app)
	message := message.NewMessageClient(app)
	school := school.NewSchoolClient(app)
	moment := moment.NewMomentClient(app)
	messageTemplate := messageTemplate.NewMessageTemplate(app)
	groupChat := groupChat.NewGroupChat(app)

	return client,
		contactWayClient,
		statistics,
		message,
		school,
		moment,
		messageTemplate,
		groupChat
}
