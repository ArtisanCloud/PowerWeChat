package miniProgram

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/auth"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/base"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/customerServiceMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/dataCube"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/express"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/image"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/immediateDelivery"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/internet"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/liveBroadcast"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/nearbyPoi"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/ocr"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/operation"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/plugin"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/riskControl"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/search"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/security"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/serviceMarket"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/shortLink"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/soter"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/uniformMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/updatableMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/urlLink"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/urlScheme"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/wxaCode"
	"net/http"
)

type MiniProgram struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	Auth        *auth.Client

	ActiveMessage *updatableMessage.Client

	Broadcast              *liveBroadcast.Client
	CustomerServiceMessage *customerServiceMessage.Client

	DataCube       *dataCube.Client
	UniformMessage *uniformMessage.Client

	Image *image.Client

	Internet *internet.Client

	Express  *express.Client
	Delivery *immediateDelivery.Client

	OCR       *ocr.Client
	Operation *operation.Client
	Plugin    *plugin.Client

	NearbyPoi *nearbyPoi.Client

	WXACode *wxaCode.Client

	URLScheme *urlScheme.Client
	URLLink   *urlLink.Client

	Security *security.Client
	Search   *search.Client

	ServiceMarket *serviceMarket.Client

	SubscribeMessage *subscribeMessage.Client

	ShortLink *shortLink.Client
	Soter     *soter.Client

	UpdatableMessage *updatableMessage.Client

	RiskControl *riskControl.Client

	Config *kernel.Config
}

type UserConfig struct {
	AppID  string
	Secret string

	ResponseType string
	Log          Log

	HttpDebug bool
	Debug     bool
}

type Log struct {
	Level string
	File  string
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
				"base_uri": "https://api.weixin.qq.com/",
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
	app.AccessToken = auth.RegisterProvider(app)
	app.Auth = auth.RegisterAuthProvider(app)

	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register Broadcast --------------
	app.Broadcast = liveBroadcast.RegisterProvider(app)

	//-------------- register CustomerServiceMessage --------------
	app.CustomerServiceMessage = customerServiceMessage.RegisterProvider(app)

	//-------------- register Data cube --------------
	app.DataCube = dataCube.RegisterProvider(app)

	//-------------- register ActiveMessage --------------
	app.ActiveMessage = updatableMessage.RegisterProvider(app)

	//-------------- register Message --------------
	app.UniformMessage = uniformMessage.RegisterProvider(app)

	//-------------- register Image --------------
	app.Image = image.RegisterProvider(app)

	//-------------- register Internet --------------
	app.Internet = internet.RegisterProvider(app)

	//-------------- register Express --------------
	app.Express = express.RegisterProvider(app)

	//-------------- register Delivery --------------
	app.Delivery = immediateDelivery.RegisterProvider(app)

	//-------------- register OCR --------------
	app.OCR = ocr.RegisterProvider(app)
	//-------------- register Operation --------------
	app.Operation = operation.RegisterProvider(app)

	//-------------- register Plugin --------------
	app.Plugin = plugin.RegisterProvider(app)

	//-------------- register NearbyPoi --------------
	app.NearbyPoi = nearbyPoi.RegisterProvider(app)

	//-------------- register WXACode --------------
	app.WXACode = wxaCode.RegisterProvider(app)

	//-------------- register URLScheme --------------
	app.URLScheme = urlScheme.RegisterProvider(app)
	//-------------- register URLLink --------------
	app.URLLink = urlLink.RegisterProvider(app)

	//-------------- register Security --------------
	app.Security = security.RegisterProvider(app)

	//-------------- register Search --------------
	app.Search = search.RegisterProvider(app)

	//-------------- register ShortLink --------------
	app.ShortLink = shortLink.RegisterProvider(app)

	//-------------- register Soter --------------
	app.Soter = soter.RegisterProvider(app)

	//-------------- register Service Market --------------
	app.ServiceMarket = serviceMarket.RegisterProvider(app)

	//-------------- register SubscribeMessage --------------
	app.SubscribeMessage = subscribeMessage.RegisterProvider(app)

	//-------------- register UpdatableMessage --------------
	app.UpdatableMessage = updatableMessage.RegisterProvider(app)

	//-------------- register RiskControl --------------
	app.RiskControl = riskControl.RegisterProvider(app)

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
	case "ActiveMessage":
		return app.ActiveMessage
	case "Broadcast":
		return app.Broadcast
	case "CustomerServiceMessage":
		return app.CustomerServiceMessage

	case "DataCube":
		return app.DataCube

	case "Image":
		return app.Image
	case "Internet":
		return app.Internet

	case "UniformMessage":
		return app.UniformMessage

	case "Express":
		return app.Express
	case "Delivery":
		return app.Delivery
	case "OCR":
		return app.OCR
	case "Operation":
		return app.Operation
	case "Plugin":
		return app.Plugin

	case "NearbyPoi":
		return app.NearbyPoi

	case "WXACode":
		return app.WXACode

	case "URLScheme":
		return app.URLScheme
	case "URLLink":
		return app.URLLink

	case "Security":
		return app.Security

	case "Search":
		return app.Search

	case "ServiceMarket":
		return app.ServiceMarket

	case "ShortLink":
		return app.ShortLink

	case "Soter":
		return app.Soter

	case "SubscribeMessage":
		return app.SubscribeMessage

	case "UpdatableMessage":
		return app.UpdatableMessage
	case "RiskControl":
		return app.RiskControl

	default:
		return nil
	}

}


func (app *MiniProgram) SetExternalRequest(r *http.Request) {
	app.Base.ExternalRequest = r
}

func (app *MiniProgram) GetExternalRequest() (r *http.Request) {
	return app.Base.ExternalRequest
}


func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},

		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
