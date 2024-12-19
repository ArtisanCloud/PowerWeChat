package openPlatform

import (
	"context"
	"errors"

	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	miniProgram2 "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	officialAccount2 "github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/auth"
	auth2 "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/auth"
	miniProgram3 "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram"
	auth3 "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/auth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/officialAccount/account"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/codeTemplate"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/component"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/server"
)

type OpenPlatform struct {
	*kernel.ServiceContainer

	Base         *base.Client
	VerifyTicket *auth.VerifyTicket
	AccessToken  *auth.AccessToken
	Server       *server.Guard
	Encryptor    *kernel.Encryptor

	CodeTemplate *codeTemplate.Client
	Component    *component.Client

	Config *kernel.Config

	Logger *logger.Logger
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
	Stdout bool
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

	// init app
	app := &OpenPlatform{
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
		"stdout":     app.Config.GetBool("log.stdout", false),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth --------------
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

	//-------------- register CodeTemplate --------------
	app.CodeTemplate, err = codeTemplate.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Component --------------
	app.Component, err = component.RegisterProvider(app)
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
	case "VerifyTicket":
		return app.VerifyTicket
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

	case "CodeTemplate":
		return app.CodeTemplate
	case "Component":
		return app.Component

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	baseURI := "https://api.weixin.qq.com/"
	if userConfig.Http.BaseURI != "" {
		baseURI = userConfig.Http.BaseURI
	}
	timeout := 5.0
	if userConfig.Http.Timeout > 0 {
		timeout = userConfig.Http.Timeout
	}
	config := &object.HashMap{

		"app_id":    userConfig.AppID,
		"secret":    userConfig.Secret,
		"auth_code": userConfig.AuthCode,
		"token":     userConfig.Token,
		"aes_key":   userConfig.AESKey,

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
			"stdout": userConfig.Log.Stdout,
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

func (app *OpenPlatform) OfficialAccount(appID string, refreshToken string, accessToken *auth2.AccessToken) (application *officialAccount.Application, err error) {

	// 配置公众号的授权者的配置
	userConfig, err := app.GetOfficialAuthorizerConfig(appID, refreshToken)
	if err != nil {
		return nil, err
	}
	// 先初始化一个OfficialAccount的账号
	application, err = officialAccount.NewApplication(userConfig)
	if err != nil {
		return nil, err
	}

	// 重塑一个AccessToken
	if accessToken == nil {
		accessToken, err = auth2.NewAccessToken(application.OfficialAccount, app)
		if err != nil {
			return nil, err
		}
	}
	application.AccessToken.AccessToken = accessToken.AccessToken
	application.Encryptor = app.Encryptor
	application.Account, err = account.NewClient(application, app)

	token, err := app.AccessToken.GetToken(context.Background(), false)
	if err != nil {
		return nil, err
	}
	application.OAuth.WithComponent(&object.HashMap{
		"id":    app.Config.GetString("app_id", ""),
		"token": token.ComponentAccessToken,
	})

	return application, err

}

func (app *OpenPlatform) MiniProgram(appID string, refreshToken string, accessToken *auth2.AccessToken) (application *miniProgram3.Application, err error) {
	userConfig, err := app.GetMiniProgramAuthorizerConfig(appID, refreshToken)
	application, err = miniProgram3.NewApplication(userConfig)

	// 重塑一个AccessToken
	if accessToken == nil {
		accessToken, err = auth2.NewAccessToken(application.MiniProgram, app)
		if err != nil {
			return nil, err
		}
	}
	application.AccessToken.AccessToken = accessToken.AccessToken
	application.Encryptor, err = miniProgram2.NewEncryptor(
		app.Config.GetString("app_id", ""),
		app.Config.GetString("token", ""),
		app.Config.GetString("aes_key", ""),
	)
	application.Auth, err = auth3.NewClient(application, app)

	return application, err

}

func (app *OpenPlatform) GetOfficialAuthorizerConfig(appID string, refreshToken string) (userConfig *officialAccount2.UserConfig, err error) {

	// 先从缓存中获取token，需要在之前的授权流程中，通过ComponentVerifyTicket生成的token。
	token, _ := app.AccessToken.GetToken(context.Background(), false)
	config := app.GetConfig()
	cacheHandle := config.Get("cache", nil).(cache.CacheInterface)

	// 将oauth的配置对象化
	var oauth = officialAccount2.OAuth{}
	err = object.HashMapToStructure(config.Get("oauth", nil).(*object.HashMap), &oauth)
	if err != nil {
		return nil, err
	}

	log := config.Get("log", nil).(*object.HashMap)

	// 用户的配置信息重新映射到UserConfig配置对象
	userConfig = &officialAccount2.UserConfig{
		ComponentAppID:    config.GetString("app_id", ""),
		ComponentAppToken: token.ComponentAccessToken,
		AppID:             appID,
		RefreshToken:      refreshToken,
		Secret:            config.GetString("secret", ""),
		Token:             config.GetString("token", ""),
		AESKey:            config.GetString("aes_key", ""),

		ResponseType: config.GetString("response_type", ""),
		Log: officialAccount2.Log{
			Level: (*log)["level"].(string),
			File:  (*log)["file"].(string),
			ENV:   (*log)["env"].(string),
		},
		OAuth: oauth,
		Cache: cacheHandle,

		HttpDebug: config.GetBool("http_debug", false),
		Debug:     config.GetBool("debug", false),
		NotifyURL: config.GetString("notify_url", ""),
		Sandbox:   config.GetBool("sandbox", false),
	}

	return userConfig, err
}

func (app *OpenPlatform) GetMiniProgramAuthorizerConfig(appID string, refreshToken string) (userConfig *miniProgram2.UserConfig, err error) {

	token, _ := app.AccessToken.GetToken(context.Background(), false)
	config := app.GetConfig()
	cache := config.Get("cache", nil).(cache.CacheInterface)
	var oauth = miniProgram2.OAuth{}
	err = object.HashMapToStructure(config.Get("oauth", nil).(*object.HashMap), &oauth)
	if err != nil {
		return nil, err
	}
	log := config.Get("log", nil).(*object.HashMap)

	userConfig = &miniProgram2.UserConfig{
		AppID:  appID,
		Secret: config.GetString("secret", ""),

		RefreshToken:      refreshToken,
		ComponentAppID:    config.GetString("app_id", ""),
		ComponentAppToken: token.ComponentAccessToken,

		ResponseType: config.GetString("secret", ""),
		Log: miniProgram2.Log{
			Level: (*log)["level"].(string),
			File:  (*log)["file"].(string),
			ENV:   (*log)["env"].(string),
		},
		OAuth: oauth,
		Cache: cache,

		HttpDebug: config.GetBool("http_debug", false),
		Debug:     config.GetBool("debug", false),
		NotifyURL: config.GetString("notify_url", ""),
		Sandbox:   config.GetBool("sandbox", false),
	}

	return userConfig, err
}

// Return the pre-authorization login page url.
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Authorization_Process_Technical_Description.html
func (app *OpenPlatform) GetFastRegistrationURL(ctx context.Context, callbackUrl string, optional *object.StringMap) (url string, err error) {

	config := app.GetConfig()

	code, err := app.Base.CreatePreAuthorizationCode(ctx)
	if err != nil {
		return url, err
	}
	if code.ErrCode != 0 {
		return url, errors.New(code.ErrMsg)
	}
	(*optional)["pre_auth_code"] = code.PreAuthCode
	queries := object.MergeStringMap(optional, &object.StringMap{
		"component_appid": config.GetString("app_id", ""),
		"redirect_uri":    callbackUrl,
	})

	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" + object.GetJoinedWithKSort(queries), err

}

// Return the pre-authorization login page url (mobile).
func (app *OpenPlatform) GetMobilePreAuthorizationURL(ctx context.Context, callbackUrl string, optional *object.StringMap) (url string, err error) {

	config := app.GetConfig()

	code, err := app.Base.CreatePreAuthorizationCode(ctx)
	if err != nil {
		return url, err
	}
	if code.ErrCode != 0 {
		return url, errors.New(code.ErrMsg)
	}
	(*optional)["pre_auth_code"] = code.PreAuthCode
	queries := object.MergeStringMap(
		&object.StringMap{
			"auth_type": "3",
		},
		optional,
		&object.StringMap{
			"component_appid": config.GetString("app_id", ""),
			"redirect_uri":    callbackUrl,
			"action":          "bindcomponent",
			"no_scan":         "1",
		})

	return "https://mp.weixin.qq.com/safe/bindcomponent?" + object.GetJoinedWithKSort(queries) + "#wechat_redirect", err

}
