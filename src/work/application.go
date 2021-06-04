package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type Work struct {
	Container *kernel.ServiceContainer
	Base interface{}
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
	}

	// inject components
	//auth.startInit()
	//base.startInit()
	//base.startInit()
	//auth.startInit()
	//base.startInit()
	//base.startInit()
	//auth3.intall()
	//base3.startInit()
	//base.startInit()
	//auth.startInit()
	//base.startInit()
	//base.startInit()

	return app
}


