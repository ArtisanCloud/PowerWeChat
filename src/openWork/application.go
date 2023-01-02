package openWork

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	miniProgram2 "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/auth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/contact"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/corp"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/device"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/media"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/provider"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server"
	suit "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
	workMiniProgram "github.com/ArtisanCloud/PowerWeChat/v3/src/work/miniProgram"
)

type OpenWork struct {
	*kernel.ServiceContainer

	Base                *base.Client
	Server              *server.Guard
	Corp                *corp.Client
	Provider            *provider.Client
	SuiteAccessToken    *suit.AccessToken
	ProviderAccessToken *auth.AccessToken
	SuiteTicket         *suit.SuiteTicket
	MiniProgram         *miniProgram.Client
	Media               *media.Client
	Contact             *contact.Client
	LicenseOrder        *license.Client
	LicenseAccount      *license.Account
	Device              *device.Client

	Encryptor *kernel.Encryptor
	Config    *kernel.Config
	Logger    *logger.Logger
}

type UserConfig struct {
	AppID    string
	Secret   string
	AuthCode string
	Token    string
	AESKey   string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
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
				"timeout":  5.0,
				"base_uri": "https://api.weixin.qq.com/",
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

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat.log"),
		"errorPath":  app.Config.GetString("log.file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth --------------
	app.SuiteTicket, app.SuiteAccessToken, err = suit.RegisterProvider(app)
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

	//-------------- register Corp --------------
	app.Corp, err = corp.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Provider --------------
	app.Provider, err = provider.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register ProviderAccessToken --------------
	app.ProviderAccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Media --------------
	app.Media, err = media.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Contact --------------
	app.Contact, err = contact.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register License --------------
	app.LicenseOrder, app.LicenseAccount, err = license.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Device --------------
	app.Device, err = device.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *OpenWork) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OpenWork) GetAccessToken() *kernel.AccessToken {
	return app.ProviderAccessToken.AccessToken
}

func (app *OpenWork) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OpenWork) GetComponent(name string) interface{} {

	switch name {
	case "Corp":
		return app.Corp
	case "Provider":
		return app.Provider
	case "SuiteAccessToken":
		return app.SuiteAccessToken
	case "ProviderAccessToken":
		return app.ProviderAccessToken
	case "SuiteTicket":
		return app.SuiteTicket
	case "MiniProgram":
		return app.MiniProgram
	case "Media":
		return app.Media
	case "Contact":
		return app.Contact
	case "LicenseOrder":
		return app.LicenseOrder
	case "LicenseAccount":
		return app.LicenseAccount
	case "Device":
		return app.Device

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id":    userConfig.AppID,
		"secret":    userConfig.Secret,
		"auth_code": userConfig.AuthCode,
		"token":     userConfig.Token,
		"aes_key":   userConfig.AESKey,

		"response_type": userConfig.ResponseType,
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
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

func (app *OpenWork) GetMiniProgram() (application *workMiniProgram.Application, err error) {

	config := app.GetConfig()
	userConfig, err := app.convertUserConfigToWorkMiniProgramConfig(config.Collection)

	application, err = workMiniProgram.NewApplication(userConfig)

	return application, err

}

func (app *OpenWork) convertUserConfigToWorkMiniProgramConfig(config *object.Collection) (userConfig *workMiniProgram.UserConfig, err error) {

	// log config
	logConfig := config.Get("log", nil)
	log := miniProgram2.Log{
		Level: "",
		File:  "",
		ENV:   "",
	}
	if logConfig != nil {
		log = logConfig.(miniProgram2.Log)
	}

	//// auth config
	//logConfig := config.Get("log", nil)
	//log := miniProgram2.Log{
	//	Level: "",
	//	File:  "",
	//	ENV:   "",
	//}
	//if logConfig != nil {
	//	log = logConfig.(miniProgram2.Log)
	//}
	//
	//// cache config
	//cacheConfig := config.Get("log", nil)
	//var cache cache.CacheInterface
	//if cacheConfig != nil {
	//	cache = cacheConfig.(cache.CacheInterface)
	//}

	// http config
	httpConfig := config.Get("http", nil)
	var mapHttpConfig *object.HashMap
	if httpConfig != nil {
		mapHttpConfig = httpConfig.(*object.HashMap)
	}
	userConfig = &workMiniProgram.UserConfig{

		MiniProgramUserConfig: &miniProgram2.UserConfig{
			AppID:             config.GetString("appID", ""),
			Secret:            config.GetString("secret", ""),
			RefreshToken:      config.GetString("refreshToken", ""),
			ComponentAppID:    config.GetString("componentAppID", ""),
			ComponentAppToken: config.GetString("componentAppToken", ""),
			ResponseType:      config.GetString("responseType", ""),
			Log:               log,
			//OAuth:             config.GetString("oauth", ""),
			//Cache:             cache,
			HttpDebug: config.GetBool("httpDebug", false),
			Debug:     config.GetBool("debug", false),
			NotifyURL: config.GetString("notifyURL", ""),
			Sandbox:   config.GetBool("sandbox", false),
		},
		CorpID:      config.GetString("corpID", ""),
		AgentID:     config.GetInt("agentID", 0),
		Secret:      config.GetString("secret", ""),
		Token:       config.GetString("token", ""),
		AESKey:      config.GetString("aESKey", ""),
		CallbackURL: config.GetString("callbackURL", ""),
		Http:        mapHttpConfig,
	}

	return userConfig, err
}
