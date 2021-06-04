package kernel

import "github.com/ArtisanCloud/go-libs/object"

type ApplicationInterface interface {
	GetConfig() object.HashMap
}

type ServiceContainer struct {
	ID int

	DefaultConfig object.HashMap
	UserConfig    object.HashMap
}

func (container *ServiceContainer) getBaseConfig() object.HashMap {
	return object.HashMap{
		// http://docs.guzzlephp.org/en/stable/request-options.html
		"http": object.HashMap{
			"timeout":  30.0,
			"base_uri": "https://api.weixin.qq.com/",
		},
	}
}

func (container *ServiceContainer) GetConfig() []object.HashMap {
	baseConfig := container.getBaseConfig()

	return []object.HashMap{
		baseConfig,
		container.DefaultConfig,
		container.UserConfig,
	}
}
