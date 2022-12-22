package kernel

import (
	"crypto/md5"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type ApplicationInterface interface {
	GetContainer() *ServiceContainer
	GetAccessToken() *AccessToken
	GetConfig() *Config
	GetComponent(name string) interface{}
}

type ExtraInfo struct {
	Prepends *power.HashMap
	ID       string
}

func NewExtraInfo() (*ExtraInfo, error) {
	return &ExtraInfo{
		Prepends: &power.HashMap{},
		ID:       "",
	}, nil
}

type ServiceContainer struct {
	ID string

	Prepends      *object.Attribute
	DefaultConfig *object.HashMap
	UserConfig    *object.HashMap
	Config        *object.HashMap
}

func NewServiceContainer(config *object.HashMap, extraInfos ...*ExtraInfo) (*ServiceContainer, error) {

	var extraInfo = &ExtraInfo{}
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	prepends, err := power.PowerHashMapToObjectHashMap(extraInfo.Prepends)
	if err != nil {
		return nil, err
	}
	container := &ServiceContainer{
		ID:         extraInfo.ID,
		UserConfig: config,
		Prepends:   object.NewAttribute(prepends),
	}

	return container, nil
}

func (container *ServiceContainer) GetID() string {

	if container.ID == "" {
		data, _ := object.JsonEncode(container.UserConfig)
		encodeData := md5.Sum([]byte(data))
		container.ID = string(encodeData[:])
	}

	return container.ID
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
	basicConfig := container.getBaseConfig()

	// merge config
	container.Config = object.ReplaceHashMapRecursive(container.Config, basicConfig, container.DefaultConfig, container.UserConfig)
	//fmt.Dump(container.Config)
	return container.Config
}
