package openWork

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/externalcontact"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/provider"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server"
	suit "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/user"
)

type OpenWork struct {
	*kernel.ServiceContainer

	Base             *base.Client
	Server           *server.Guard
	User             *user.Client
	ExternalContact  *externalcontact.Client
	Suite            *suit.Client
	SuiteAccessToken *suit.AccessToken
	SuiteTicket      *suit.SuiteTicket

	Encryptor *kernel.Encryptor
	Config    *kernel.Config
	Logger    *logger.Logger
}

type UserConfig struct {
	AppID          string
	Secret         string
	ProviderCorpID string
	ProviderSecret string
	AuthCode       string
	Token          string
	AESKey         string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	Http Http

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
}

type Http struct {
	Timeout  float64
	BaseURI  string
	ProxyURI string
}

type Log struct {
	Driver contract.LoggerInterface
	Level  string
	File   string
	Error  string
	ENV    string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewOpenWork(config *UserConfig) (*OpenWork, error) {
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
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}

	// init app
	app := &OpenWork{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger(app.Config.Get("log.driver", nil), &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat/info.log"),
		"errorPath":  app.Config.GetString("log.error", "./wechat/error.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth --------------
	app.SuiteTicket, app.SuiteAccessToken, err = suit.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	app.Suite, err = suit.NewClient(app)
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

	//-------------- register User --------------
	app.User, err = user.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *OpenWork) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OpenWork) GetAccessToken() *kernel.AccessToken {
	return app.SuiteAccessToken.AccessToken
}

func (app *OpenWork) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OpenWork) Provider(corpID string) (*provider.Client, error) {
	return provider.NewClient(app, corpID)
}

func (app *OpenWork) License(corpID string) (*license.Client, error) {
	return license.NewClient(app, corpID)
}

func (app *OpenWork) GetComponent(name string) interface{} {

	switch name {
	case "User":
		return app.User
	case "ExternalContact":
		return app.ExternalContact
	case "Suite":
		return app.Suite
	case "SuiteAccessToken":
		return app.SuiteAccessToken
	case "SuiteTicket":
		return app.SuiteTicket
	case "Encryptor":
		return app.Encryptor
	case "Server":
		return app.Server
	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	baseURI := "https://qyapi.weixin.qq.com/"
	if userConfig.Http.BaseURI != "" {
		baseURI = userConfig.Http.BaseURI
	}
	timeout := 5.0
	if userConfig.Http.Timeout > 0 {
		timeout = userConfig.Http.Timeout
	}
	config := &object.HashMap{

		"app_id":          userConfig.AppID,
		"secret":          userConfig.Secret,
		"provider_corpid": userConfig.ProviderCorpID,
		"provider_secret": userConfig.ProviderSecret,
		"auth_code":       userConfig.AuthCode,
		"token":           userConfig.Token,
		"aes_key":         userConfig.AESKey,

		"response_type": userConfig.ResponseType,
		"http": &object.HashMap{
			"timeout":   timeout,
			"base_uri":  baseURI,
			"proxy_uri": userConfig.Http.ProxyURI,
		},
		"log": &object.HashMap{
			"driver": userConfig.Log.Driver,
			"level":  userConfig.Log.Level,
			"file":   userConfig.Log.File,
			"error":  userConfig.Log.Error,
			"env":    userConfig.Log.ENV,
		},
		"cache":      userConfig.Cache,
		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,

		"oauth": &object.HashMap{
			"scopes":    userConfig.OAuth.Scopes,
			"callbacks": userConfig.OAuth.Callback,
		},
	}

	return config, nil

}
