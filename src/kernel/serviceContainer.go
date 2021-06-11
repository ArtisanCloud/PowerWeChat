package kernel

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ApplicationInterface interface {
	GetContainer() *ServiceContainer
	GetAccessToken() *AccessToken
}

type ServiceContainer struct {
	ID int

	DefaultConfig *object.HashMap
	UserConfig    *object.HashMap
	Config        *object.HashMap
}

func (container *ServiceContainer) getBaseConfig() *object.HashMap {
	return &object.HashMap{
		// http://docs.guzzlephp.org/en/stable/request-options.html
		"http": object.HashMap{
			"timeout":  30.0,
			"base_uri": "https://api.weixin.qq.com/",
		},
	}
}

func (container *ServiceContainer) GetConfig() *object.HashMap {

	// init container config
	var config *object.HashMap
	if container.Config != nil {
		config = container.Config
	} else {
		config = container.getBaseConfig()
	}

	// merge config
	container.Config = object.MergeHashMap(config, container.DefaultConfig, container.UserConfig)

	return container.Config
}
