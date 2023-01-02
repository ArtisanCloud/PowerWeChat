package miniProgram

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/auth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/customerServiceMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/dataCube"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/express"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/image"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/immediateDelivery"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/internet"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/liveBroadcast"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/nearbyPoi"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/ocr"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/operation"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/phoneNumber"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/plugin"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/riskControl"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/search"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/security"
	server2 "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/server"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/serviceMarket"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/shortLink"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/soter"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/uniformMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/updatableMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlScheme"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/wxaCode"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server"
)

type MiniProgram struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	Auth        *auth.Client

	Server *server.Guard

	ActiveMessage *updatableMessage.Client

	Encryptor *Encryptor

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

	PhoneNumber *phoneNumber.Client

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

	Logger *logger.Logger
}

type UserConfig struct {
	AppID  string
	Secret string

	RefreshToken      string
	ComponentAppID    string
	ComponentAppToken string
	Token             string
	AESKey            string

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

func NewMiniProgram(config *UserConfig, extraInfos ...*kernel.ExtraInfo) (*MiniProgram, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	// init an app container
	container, err := kernel.NewServiceContainer(userConfig, extraInfo)
	if err != nil {
		return nil, err
	}
	container.GetConfig()

	// init app
	app := &MiniProgram{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("log.env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat.log"),
		"errorPath":  app.Config.GetString("log.file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth,AccessToken --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	app.Auth, err = auth.RegisterAuthProvider(app)
	if err != nil {
		return nil, err
	}

	app.PhoneNumber, err = phoneNumber.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	// -------------- register Encryptor --------------
	app.Encryptor, err = NewEncryptor(
		(app.Config).GetString("app_id", ""),
		(app.Config).GetString("token", ""),
		(app.Config).GetString("aes_key", ""),
	)
	if err != nil {
		return nil, err
	}

	//-------------- register Encryptor and Server --------------
	app.Server, err = server2.RegisterProvider(app)

	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Broadcast --------------
	app.Broadcast, err = liveBroadcast.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register CustomerServiceMessage --------------
	app.CustomerServiceMessage, err = customerServiceMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Data cube --------------
	app.DataCube, err = dataCube.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register ActiveMessage --------------
	app.ActiveMessage, err = updatableMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Message --------------
	app.UniformMessage, err = uniformMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Image --------------
	app.Image, err = image.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Internet --------------
	app.Internet, err = internet.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Express --------------
	app.Express, err = express.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Delivery --------------
	app.Delivery, err = immediateDelivery.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register OCR --------------
	app.OCR, err = ocr.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Operation --------------
	app.Operation, err = operation.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Plugin --------------
	app.Plugin, err = plugin.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register NearbyPoi --------------
	app.NearbyPoi, err = nearbyPoi.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register WXACode --------------
	app.WXACode, err = wxaCode.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register URLScheme --------------
	app.URLScheme, err = urlScheme.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register URLLink --------------
	app.URLLink, err = urlLink.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Security --------------
	app.Security, err = security.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Search --------------
	app.Search, err = search.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register ShortLink --------------
	app.ShortLink, err = shortLink.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Soter --------------
	app.Soter, err = soter.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Service Market --------------
	app.ServiceMarket, err = serviceMarket.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register SubscribeMessage --------------
	app.SubscribeMessage, err = subscribeMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register UpdatableMessage --------------
	app.UpdatableMessage, err = updatableMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register RiskControl --------------
	app.RiskControl, err = riskControl.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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
	case "auth":
		return app.Auth
	case "Config":
		return app.Config
	case "Server":
		return app.Server
	case "Encryptor":
		return app.Encryptor.BaseEncryptor

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

		"token":               userConfig.Token,
		"aes_key":             userConfig.AESKey,
		"refresh_token":       userConfig.RefreshToken,
		"component_app_id":    userConfig.ComponentAppID,
		"component_app_token": userConfig.ComponentAppToken,

		"response_type": userConfig.ResponseType,
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
		},
		"cache": userConfig.Cache,

		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
