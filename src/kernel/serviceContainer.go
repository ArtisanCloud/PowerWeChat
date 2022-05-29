package kernel

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
)

type ApplicationInterface interface {
	GetContainer() *ServiceContainer
	GetAccessToken() *AccessToken
	GetConfig() *Config
	GetComponent(name string) interface{}
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
		"http": &object.HashMap{
			"timeout":  30.0,
			"base_uri": "https://api.weixin.qq.com/",
		},
	}
}

func (container *ServiceContainer) GetConfig() *object.HashMap {

	// init container config
	var config *object.HashMap

	basicConfig := container.getBaseConfig()

	// merge config
	container.Config = object.ReplaceHashMapRecursive(config, basicConfig, container.DefaultConfig, container.UserConfig)
	//fmt.Dump(container.Config)
	return container.Config
}
