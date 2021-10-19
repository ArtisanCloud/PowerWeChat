package externalContact

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/groupWelcomeTemplate"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/moment"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/momentStrategy"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/school"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/statistics"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/tag"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/transfer"
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
