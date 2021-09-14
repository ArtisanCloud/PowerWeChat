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
	*contactWay.Client,
	*statistics.Client,
	*message.Client,
	*school.Client,
	*moment.Client,
	*messageTemplate.Client,
	*groupChat.Client,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	contactWayClient := contactWay.NewClient(app)

	statistics := statistics.NewClient(app)
	message := message.NewClient(app)
	school := school.NewClient(app)
	moment := moment.NewClient(app)
	messageTemplate := messageTemplate.NewClient(app)
	groupChat := groupChat.NewClient(app)

	return client,
		contactWayClient,
		statistics,
		message,
		school,
		moment,
		messageTemplate,
		groupChat
}
