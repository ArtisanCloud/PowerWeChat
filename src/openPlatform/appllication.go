package openPlatform

import (
	"github.com/ArtisanCloud/PowerLibs/v2/logger"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/auth"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/base"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/server"
)

type OpenPlatform struct {
	*kernel.ServiceContainer

	Base         *base.Client
	VerifyTicket *auth.VerifyTicket
	AccessToken  *auth.AccessToken
	Server       *server.Guard
	Encryptor    *kernel.Encryptor

	Config *kernel.Config

	Logger *logger.Logger
}

type UserConfig struct {
	AppID  string
	Secret string
	//Token  string
	//AESKey string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	HttpDebug bool
	Debug     bool
}

type Log struct {
	Level string
	File  string
	ENV   string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewOpenPlatform(config *UserConfig) (*OpenPlatform, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": &object.HashMap{
				"timeout":  5.0,
				"base_uri": "https://api.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &OpenPlatform{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.VerifyTicket, app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Encryptor and Server --------------
	app.Encryptor, app.Server, err = server.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("file", "./wechat.log"),
		"errorPath":  app.Config.GetString("file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *OpenPlatform) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OpenPlatform) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *OpenPlatform) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OpenPlatform) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "Config":
		return app.Config

	case "Server":
		return app.Server
	case "Encryptor":
		return app.Encryptor

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,
		//"token":   userConfig.Token,
		//"aes_key": userConfig.AESKey,

		"response_type": userConfig.ResponseType,
		"log": &object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
		},
		"cache":      userConfig.Cache,
		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,

		"oauth": &object.HashMap{
			"scopes":   userConfig.OAuth.Scopes,
			"callback": userConfig.OAuth.Callback,
		},
	}

	return config, nil

}
