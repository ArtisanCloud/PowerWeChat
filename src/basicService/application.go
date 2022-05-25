package officialAccount

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/contentSecurity"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/media"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/qrCode"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/url"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/providers"
	"net/http"
)

type Application struct {
	*kernel.ServiceContainer

	Base *kernel.BaseClient

	Config *kernel.Config

	JSSDK           *jssdk.Client
	QRCode          *qrCode.Client
	Media           *media.Client
	URL             *url.Client
	ContentSecurity *contentSecurity.Client
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

func NewApplication(config *UserConfig) (*Application, error) {
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
	app := &Application{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register JSSDK --------------
	app.JSSDK = jssdk.RegisterProvider(app)
	//-------------- register QRCode --------------
	app.QRCode = qrCode.RegisterProvider(app)

	//-------------- register Media --------------
	app.Media = media.RegisterProvider(app)

	//-------------- media --------------
	app.URL = url.RegisterProvider(app)
	app.ContentSecurity = contentSecurity.RegisterProvider(app)

	return app, err
}

func (app *Application) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Application) GetAccessToken() *kernel.AccessToken {
	return nil
}

func (app *Application) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Application) GetComponent(name string) interface{} {

	switch name {
	case "JSSDK":
		return app.JSSDK
	case "QRCode":
		return app.QRCode
	case "Media":
		return app.Media
	case "URL":
		return app.URL
	case "ContentSecurity":
		return app.ContentSecurity

	default:
		return nil
	}

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
