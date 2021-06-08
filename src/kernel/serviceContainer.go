package kernel

import "github.com/ArtisanCloud/go-libs/object"



type ApplicationInterface interface {

	GetComponents() *object.HashMap
	GetContainer() *ServiceContainer
	GetApp() ApplicationInterface
}



type ServiceContainer struct {
	ID int

	DefaultConfig *object.HashMap
	UserConfig    *object.HashMap
	Config        *object.HashMap
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

func (container *ServiceContainer) GetConfig() *object.HashMap {
	// if container config has been setup
	if container.Config!=nil{
		return container.Config
	}

	container.Config = &object.HashMap{}

	// merge config
	container.Config = object.MergeHashMap(container.DefaultConfig,container.Config)
	container.Config = object.MergeHashMap(container.UserConfig,container.Config)

	return container.Config
}
