package officialAccount

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/work/auth"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"github.com/ArtisanCloud/power-wechat/src/work/media"
	"github.com/ArtisanCloud/power-wechat/src/work/menu"
	"github.com/ArtisanCloud/power-wechat/src/work/oauth"
	"net/http"
)

type OfficialAccount struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config *kernel.Config
	
	Media *media.Client
	Menu  *menu.Client
}

type UserConfig struct {
	AppID           string
	Secret           string

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

func NewOfficialAccount(config *UserConfig) (*OfficialAccount, error) {
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
	app := &OfficialAccount{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.AccessToken = auth.RegisterProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

	//-------------- menu --------------
	app.Menu = menu.RegisterProvider(app)


	return app, err
}

func (app *OfficialAccount) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OfficialAccount) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *OfficialAccount) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OfficialAccount) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "OAuth":
		return app.OAuth
	case "Config":
		return app.Config


	case "Menu":
		return app.Menu


	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id":           userConfig.AppID,
		"secret":             userConfig.Secret,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
	}

	return config, nil

}
