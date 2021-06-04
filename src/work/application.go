package work

import "github.com/ArtisanCloud/go-libs/object"

type Work struct {
	DefaultConfig object.HashMap
	UserConfig    interface{}

	id int
}

func NewWork(config interface{}) *Work {
	app := &Work{
		UserConfig: config,
	}

	app.DefaultConfig = object.HashMap{
		"http": map[string]string{
			"base_uri": "https://qyapi.weixin.qq.com/",
		},
	}




	return app
}
