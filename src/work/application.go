package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/work/auth"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"github.com/ArtisanCloud/power-wechat/src/work/department"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact"
	"github.com/ArtisanCloud/power-wechat/src/work/media"
	"github.com/ArtisanCloud/power-wechat/src/work/message"
	"github.com/ArtisanCloud/power-wechat/src/work/oauth"
	"github.com/ArtisanCloud/power-wechat/src/work/server"
	"github.com/ArtisanCloud/power-wechat/src/work/user"
	"net/http"
)

type Work struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager
	Config      *kernel.Config
	Department  *department.Client

	Message  *message.Client
	Messager *message.Messager

	Encryptor *kernel.Encryptor
	Server    *server.Guard

	UserClient           *user.Client
	UserBatchJobsClient  *user.BatchJobsClient
	UserLinkedCorpClient *user.LinkedCorpClient
	UserTagClient        *user.TagClient

	ExternalContact                *externalContact.Client
	ExternalContactContactWay      *externalContact.ContactWayClient
	ExternalContactStatistics      *externalContact.StatisticsClient
	ExternalContactMessage         *externalContact.MessageClient
	ExternalContactSchool          *externalContact.SchoolClient
	ExternalContactMoment          *externalContact.MomentClient
	ExternalContactMessageTemplate *externalContact.MessageTemplateClient

	Media *media.Client
}

func NewWork(config *object.HashMap, r *http.Request) (*Work, error) {
	var err error

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: config,
		DefaultConfig: &object.HashMap{
			"http": object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	//-------------- external request --------------
	app.ExternalRequest = r
	if r == nil {
		app.ExternalRequest = &http.Request{}
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.AccessToken = auth.RegisterProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)

	//-------------- register Department --------------
	app.Department = department.RegisterProvider(app)

	//-------------- register Message --------------
	app.Message, app.Messager = message.RegisterProvider(app)

	//-------------- register Encryptor --------------
	app.Encryptor, app.Server = server.RegisterProvider(app)

	//-------------- register user --------------
	app.UserClient,
	app.UserBatchJobsClient,
	app.UserLinkedCorpClient,
	app.UserTagClient = user.RegisterProvider(app)

	//-------------- register external contact --------------
	app.ExternalContact,
		app.ExternalContactContactWay,
		app.ExternalContactStatistics,
		app.ExternalContactMessage,
		app.ExternalContactSchool,
		app.ExternalContactMoment,
		app.ExternalContactMessageTemplate = externalContact.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

	return app, err
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Work) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *Work) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Work) GetComponent(name string) interface{} {

	switch name {
	case "ExternalRequest":
		return app.ExternalRequest
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "OAuth":
		return app.OAuth
	case "Config":
		return app.Config
	case "Department":
		return app.Department

	case "Message":
		return app.Message
	case "Messager":
		return app.Messager

	case "Encryptor":
		return app.Encryptor
	case "Server":
		return app.Server

	case "UserClient":
		return app.UserClient
	case "UserBatchJobsClient":
		return app.UserBatchJobsClient
	case "UserLinkedCorpClient":
		return app.UserLinkedCorpClient
	case "UserTagClient":
		return app.UserTagClient

	case "ExternalContact":
		return app.ExternalContact
	case "ExternalContactContactWay":
		return app.ExternalContactContactWay
	case "ExternalContactStatistics":
		return app.ExternalContactStatistics
	case "ExternalContactMessage":
		return app.ExternalContactMessage
	case "ExternalContactSchool":
		return app.ExternalContactSchool
	case "ExternalContactMoment":
		return app.ExternalContactMoment
	case "ExternalContactMessageTemplate":
		return app.ExternalContactMessageTemplate

	default:
		return nil
	}

}
