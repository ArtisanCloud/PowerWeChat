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
	app := &Work{
		ServiceContainer: &kernel.ServiceContainer{
			UserConfig: config,
			DefaultConfig: &object.HashMap{
				"http": map[string]string{
					"base_uri": "https://qyapi.weixin.qq.com/",
				},
			},
		},
	}

	// inject components
	app.AccessToken = auth.RegisterProvider(app)
	app.Base = base.RegisterProvider(app)

	return app
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}
