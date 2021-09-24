package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupWelcomeTemplate"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/moment"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/momentStrategy"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/school"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/statistics"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/tag"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/transfer"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*contactWay.Client,
	*customerStrategy.Client,
	*groupChat.Client,
	*groupWelcomeTemplate.Client,
	*messageTemplate.Client,
	*moment.Client,
	*momentStrategy.Client,
	*school.Client,
	*statistics.Client,
	*tag.Client,
	*transfer.Client,
) {
	//config := app.GetConfig()

	Client := NewClient(app)
	ContactWayClient := contactWay.NewClient(app)
	CustomerStrategy := customerStrategy.NewClient(app)
	GroupChat := groupChat.NewClient(app)
	GroupWelcomeTemplate := groupWelcomeTemplate.NewClient(app)
	MessageTemplate := messageTemplate.NewClient(app)
	Moment := moment.NewClient(app)
	MomentStrategy := momentStrategy.NewClient(app)
	School := school.NewClient(app)
	Statistics := statistics.NewClient(app)
	Tag := tag.NewClient(app)
	Transfer := transfer.NewClient(app)

	return Client,
		ContactWayClient,
		CustomerStrategy,
		GroupChat,
		GroupWelcomeTemplate,
		MessageTemplate,
		Moment,
		MomentStrategy,
		School,
		Statistics,
		Tag,
		Transfer

}
