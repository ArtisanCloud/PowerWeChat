package externalContact

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupWelcomeTemplate"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/moment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/momentStrategy"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/school"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/statistics"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/transfer"
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
	error,
) {
	//config := app.GetConfig()

	Client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	ContactWayClient, err := contactWay.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	CustomerStrategy, err := customerStrategy.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	GroupChat, err := groupChat.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	GroupWelcomeTemplate, err := groupWelcomeTemplate.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	MessageTemplate, err := messageTemplate.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Moment, err := moment.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	MomentStrategy, err := momentStrategy.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	School, err := school.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Statistics, err := statistics.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Tag, err := tag.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Transfer, err := transfer.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

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
		Transfer,
		nil
}
