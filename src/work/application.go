package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/work/auth"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"github.com/ArtisanCloud/power-wechat/src/work/department"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact"
	"github.com/ArtisanCloud/power-wechat/src/work/message"
	"github.com/ArtisanCloud/power-wechat/src/work/oauth"
	"github.com/ArtisanCloud/power-wechat/src/work/server"
	"net/http"
)

type Work struct {
	*kernel.ServiceContainer

	ExternalRequest                *http.Request

	Base                           *base.Client
	AccessToken                    *auth.AccessToken
	OAuth                          *oauth.Manager
	Config                         *kernel.Config
	Department                     *department.Client

	Message                        *message.Client
	Messager                       *message.Messager

	Encryptor                      *kernel.Encryptor
	Server                         *server.Guard

	ExternalContact                *externalContact.Client
	ContactWay                     *externalContact.ContactWayClient
	ExternalContactStatistics      *externalContact.StatisticsClient
	ExternalContactMessage         *externalContact.MessageClient
	School                         *externalContact.SchoolClient
	ExternalContactMoment          *externalContact.MomentClient
	ExternalContactMessageTemplate *externalContact.MessageTemplateClient
}

func NewWork(config *object.HashMap, r *http.Request) (*Work, error) {
	var err error

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: config,
		DefaultConfig: &object.HashMap{
			"http": map[string]string{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	// register external request
	app.ExternalRequest = r
	if r == nil {
		app.ExternalRequest = &http.Request{}
	}

	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	// register Auth
	app.AccessToken = auth.RegisterProvider(app)
	// register Base
	app.Base = base.RegisterProvider(app)

	// register oauth
	app.OAuth, err = oauth.RegisterProvider(app)

	// register Department
	app.Department = department.RegisterProvider(app)

	// register Message
	app.Message, app.Messager = message.RegisterProvider(app)

	// register Encryptor
	app.Encryptor, app.Server = server.RegisterProvider(app)

	// register external contact
	app.ExternalContact,
		app.ContactWay,
		app.ExternalContactStatistics,
		app.ExternalContactMessage,
		app.School,
		app.ExternalContactMoment,
		app.ExternalContactMessageTemplate = externalContact.RegisterProvider(app)

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

	if name == "ExternalRequest" {
		return app.ExternalRequest
	} else if name == "Base" {
		return app.Base
	} else if name == "AccessToken" {
		return app.AccessToken
	} else if name == "OAuth" {
		return app.OAuth
	} else if name == "Config" {
		return app.Config
	} else if name == "Department" {
		return app.Department
	} else if name == "Message" {
		return app.Message
	} else if name == "Messager" {
		return app.Messager
	} else if name == "Encryptor" {
		return app.Encryptor
	} else if name == "Server" {
		return app.Server
	} else if name == "ExternalContact" {
		return app.ExternalContact
	} else if name == "ContactWay" {
		return app.ContactWay
	}

	return nil

}
