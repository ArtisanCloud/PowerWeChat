package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/go-wechat/src/work/auth"
	"github.com/ArtisanCloud/go-wechat/src/work/base"
	"github.com/ArtisanCloud/go-wechat/src/work/oauth"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager
	Config 		*kernel.Config
}

func NewWork(config *object.HashMap) *Work {
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
	app.OAuth = oauth.RegisterProvider(app)



	return app
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