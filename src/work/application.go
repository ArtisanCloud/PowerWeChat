package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/auth"
	"github.com/ArtisanCloud/go-wechat/src/work/base"
)

type Work struct {
	Container *kernel.ServiceContainer

	Components *object.HashMap
}

func NewWork(config object.HashMap) *Work {
	app := &Work{
		Container: &kernel.ServiceContainer{
			UserConfig: config,
			DefaultConfig: object.HashMap{
				"http": map[string]string{
					"base_uri": "https://qyapi.weixin.qq.com/",
				},
			},
		},
		Components: &object.HashMap{},
	}

	// inject components
	auth.Inject(app)
	base.Inject(app)


	return app
}

func (app *Work)GetComponents() *object.HashMap{
	return app.Components
}

func (app *Work) GetContainer() *kernel.ServiceContainer{
	return app.Container
}