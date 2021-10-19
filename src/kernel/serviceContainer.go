package kernel

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"net/http"
)

type ApplicationInterface interface {
	GetContainer() *ServiceContainer
	GetAccessToken() *AccessToken
	GetConfig() *Config
	GetComponent(name string) interface{}

	SetExternalRequest(r *http.Request)
	GetExternalRequest() (r *http.Request)
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
	if container.Config != nil {
		config = container.Config
	} else {
		config = container.getBaseConfig()
	}

	// merge config
	container.Config = object.MergeHashMap(config, container.DefaultConfig, container.UserConfig)
	//fmt.Dump(container.Config)
	return container.Config
}
