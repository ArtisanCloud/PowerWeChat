package openWork

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/provider"
)

type OpenWorkProvider struct {
	*kernel.ServiceContainer

	Base        *base.Client
	Config      *kernel.Config
	Logger      *logger.Logger
	AccessToken *provider.AccessToken

	Client  *provider.Client
	License *license.Client
}

func NewOpenWorkProvider(config *UserConfig) (*OpenWorkProvider, error) {
	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": &object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}

	// init app
	app := &OpenWorkProvider{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger(app.Config.Get("log.driver", nil), &object.HashMap{
		"level":      app.Config.GetString("log.level", "info"),
		"env":        app.Config.GetString("log.env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat/info.log"),
		"errorPath":  app.Config.GetString("log.error", "./wechat/error.log"),
	})
	if err != nil {
		return nil, err
	}
	app.AccessToken, err = provider.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	app.Client, err = provider.NewClient(app)
	if err != nil {
		return nil, err
	}
	app.License, err = license.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *OpenWorkProvider) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OpenWorkProvider) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *OpenWorkProvider) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OpenWorkProvider) GetComponent(name string) interface{} {
	switch name {
	case "Client":
		return app.Client
	case "License":
		return app.License
	case "Logger":
		return app.Logger

	default:
		return nil
	}
}
