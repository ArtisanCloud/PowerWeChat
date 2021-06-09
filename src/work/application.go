package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/auth"
	"github.com/ArtisanCloud/go-wechat/src/work/base"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
}

func NewWork(config *object.HashMap) *Work {
	// init an app container
	container:=&kernel.ServiceContainer{
		UserConfig: config,
		DefaultConfig: &object.HashMap{
			"http": map[string]string{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	// register Auth
	app.AccessToken = auth.RegisterProvider(app)
	// register Base
	app.Base = base.RegisterProvider(app)



	return app
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Work) GetAccessToken() *kernel.AccessToken{
	return app.AccessToken.AccessToken
}