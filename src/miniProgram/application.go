package miniProgram

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/auth"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/base"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/dataCube"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/express"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/immediateDelivery"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/uniformMessage"

)

type MiniProgram struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	Auth        *auth.Client

	DataCube *dataCube.Client
	UniformMessage  *uniformMessage.Client

	Express *express.Client
	Delivery *immediateDelivery.Client

	Config *kernel.Config
}

type UserConfig struct {
	CorpID           string
	AgentID          int
	Secret           string
	Token            string
	AESKey           string
	AuthCallbackHost string

	ResponseType string
	Log          Log
	OAuth        OAuth
	HttpDebug    bool
	Debug        bool
}

type Log struct {
	Level string
	File  string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewMiniProgram(config *UserConfig) (*MiniProgram, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &MiniProgram{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth,AccessToken --------------
	app.Auth, app.AccessToken = auth.RegisterProvider(app)

	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register Data cube --------------
	app.DataCube = dataCube.RegisterProvider(app)

	//-------------- register Message --------------
	app.UniformMessage = uniformMessage.RegisterProvider(app)

	//-------------- register Express --------------
	app.Express = express.RegisterProvider(app)

	//-------------- register Delivery --------------
	app.Delivery = immediateDelivery.RegisterProvider(app)




	return app, err
}

func (app *MiniProgram) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *MiniProgram) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *MiniProgram) GetConfig() *kernel.Config {
	return app.Config
}

func (app *MiniProgram) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "Auth":
		return app.Auth
	case "Config":
		return app.Config

	case "DataCube":
		return app.DataCube

	case "UniformMessage":
		return app.UniformMessage

	case "Express":
		return app.Express
	case "Delivery":
		return app.Delivery

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"corp_id":            userConfig.CorpID,
		"agent_id":           userConfig.AgentID,
		"secret":             userConfig.Secret,
		"token":              userConfig.Token,
		"aes_key":            userConfig.AESKey,
		"auth_callback_host": userConfig.AuthCallbackHost,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"oauth.callback": userConfig.OAuth.Callback,
		"oauth.scopes":   userConfig.OAuth.Scopes,
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
	}

	return config, nil

}
