package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/go-wechat/src/work/auth"
	"github.com/ArtisanCloud/go-wechat/src/work/base"
	"github.com/ArtisanCloud/go-wechat/src/work/department"
	"github.com/ArtisanCloud/go-wechat/src/work/message"
	"github.com/ArtisanCloud/go-wechat/src/work/oauth"
	"github.com/ArtisanCloud/go-wechat/src/work/server"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager
	Config      *kernel.Config
	Department  *department.Client
	Message     *message.Client
	Messager    *message.Messager
	Encryptor   *kernel.Encryptor
	server      *server.Guard
}

func NewWork(config *object.HashMap) (*Work, error) {
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
	app.Encryptor, app.server = server.RegisterProvider(app)

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
