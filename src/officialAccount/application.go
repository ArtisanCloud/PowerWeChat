package officialAccount

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerSocialite/src"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/qrCode"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/auth"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/base"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/media"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/user"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/user/tag"
	"net/http"
)

type OfficialAccount struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken

	Config *kernel.Config

	QRCode *qrCode.Client

	SubscribeMessage *subscribeMessage.Client

	OAuth *src.SocialiteManager

	User *user.Client
	Tag  *tag.Client

	Media *media.Client
}

type UserConfig struct {
	AppID  string
	Secret string

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
			"http": &object.HashMap{
				"base_uri": "https://api.weixin.qq.com/",
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

	//-------------- register QRCode --------------
	app.QRCode = qrCode.RegisterProvider(app)

	//-------------- register SubscribeMessage --------------
	app.SubscribeMessage = subscribeMessage.RegisterProvider(app)


	//-------------- register User --------------
	app.User, app.Tag = user.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

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
	case "Config":
		return app.Config
	case "QRCode":
		return app.QRCode
	case "OAuth":
		return app.OAuth
	case "SubscribeMessage":
		return app.SubscribeMessage
	case "User":
		return app.User
	case "Tag":
		return app.Tag

	default:
		return nil
	}

}

func (app *OfficialAccount) SetExternalRequest(r *http.Request) {
	app.Base.ExternalRequest = r
}

func (app *OfficialAccount) GetExternalRequest() (r *http.Request) {
	return app.Base.ExternalRequest
}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,

		"response_type": userConfig.ResponseType,
		"log": &object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"cache":      userConfig.Cache,
		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
