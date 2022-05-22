package officialAccount

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	providers2 "github.com/ArtisanCloud/PowerSocialite/src/providers"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/media"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/qrCode"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/url"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/dataCube"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/auth"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/autoReply"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/base"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/broadcasting"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/card"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/comment"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/customerService"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/customerService/session"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/device"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/goods"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/guide"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/menu"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/ocr"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/poi"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/semantic"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/server"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/store"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/templateMessage"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/user"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/user/tag"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/wifi"

	"net/http"
)

type OfficialAccount struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken

	Config *kernel.Config

	// basic services
	Media            *media.Client
	URL              *url.Client
	QRCode           *qrCode.Client
	JSSDK            *jssdk.Client
	SubscribeMessage *subscribeMessage.Client

	Server                 *server.Guard
	Encryptor              *kernel.Encryptor
	User                   *user.Client
	UserTag                *tag.Client
	Menu                   *menu.Client
	TemplateMessage        *templateMessage.Client
	Material               *material.Client
	CustomerService        *customerService.Client
	CustomerServiceSession *session.Client
	Semantic               *semantic.Client
	DataCube               *dataCube.Client
	AutoReply              *autoReply.Client
	Broadcasting           *broadcasting.Client
	Card                   *card.Client
	Device                 *device.Client
	ShakeAround            *shakeAround.ShakeAround
	POI                    *poi.Client
	Store                  *store.Client
	Comment                *comment.Client
	OCR                    *ocr.Client
	Goods                  *goods.Client
	OAuth                  *providers2.WeChat
	Wifi                   *wifi.Client
	WifiCard               *wifi.CardClient
	WifiDevice             *wifi.DeviceClient
	WifiShop               *wifi.ShopClient
	Guide                  *guide.Client
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

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)
	//-------------- register QRCode --------------
	app.QRCode = qrCode.RegisterProvider(app)
	//-------------- register URL --------------
	app.URL = url.RegisterProvider(app)
	//-------------- register JSSDK --------------
	app.JSSDK = jssdk.RegisterProvider(app)
	//-------------- register SubscribeMessage --------------
	app.SubscribeMessage = subscribeMessage.RegisterProvider(app)

	//-------------- register Menu --------------
	app.Menu = menu.RegisterProvider(app)
	//-------------- register Material --------------
	app.Material = material.RegisterProvider(app)
	//-------------- register CustomerService --------------
	app.CustomerService, app.CustomerServiceSession = customerService.RegisterProvider(app)
	//-------------- register Semantic --------------
	app.Semantic = semantic.RegisterProvider(app)
	//-------------- register Encryptor and Server --------------
	app.Encryptor, app.Server = server.RegisterProvider(app)
	//-------------- register User --------------
	app.User, app.UserTag = user.RegisterProvider(app)
	//-------------- register Menu --------------
	app.Menu = menu.RegisterProvider(app)
	//-------------- register TemplateMessage --------------
	app.TemplateMessage = templateMessage.RegisterProvider(app)
	//-------------- register Material --------------
	app.Material = material.RegisterProvider(app)
	//-------------- register CustomerService --------------
	app.CustomerService, app.CustomerServiceSession = customerService.RegisterProvider(app)
	//-------------- register Semantic --------------
	app.Semantic = semantic.RegisterProvider(app)
	//-------------- register DataCube --------------
	app.DataCube = dataCube.RegisterProvider(app)
	//-------------- register AutoReply --------------
	app.AutoReply = autoReply.RegisterProvider(app)
	//-------------- register Broadcasting --------------
	app.Broadcasting = broadcasting.RegisterProvider(app)
	//-------------- register Card --------------
	app.Card = card.RegisterProvider(app)
	//-------------- register Device --------------
	app.Device = device.RegisterProvider(app)
	//-------------- register ShakeAround --------------
	app.ShakeAround = shakeAround.RegisterProvider(app)
	//-------------- register POI --------------
	app.POI = poi.RegisterProvider(app)
	//-------------- register Store --------------
	app.Store = store.RegisterProvider(app)
	//-------------- register Comment --------------
	app.Comment = comment.RegisterProvider(app)
	//-------------- register OCR --------------
	app.OCR = ocr.RegisterProvider(app)
	//-------------- register Goods --------------
	app.Goods = goods.RegisterProvider(app)
	//-------------- register wifi --------------
	app.Wifi, app.WifiCard, app.WifiDevice, app.WifiShop = wifi.RegisterProvider(app)
	//-------------- register Guide --------------
	app.Guide = guide.RegisterProvider(app)

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

	case "Media":
		return app.Media
	case "QRCode":
		return app.QRCode
	case "SubscribeMessage":
		return app.SubscribeMessage
	case "URL":
		return app.URL

	case "Server":
		return app.Server
	case "Encryptor":
		return app.Encryptor
	case "User":
		return app.User
	case "UserTag":
		return app.UserTag
	case "Menu":
		return app.Menu
	case "TemplateMessage":
		return app.TemplateMessage
	case "Material":
		return app.Material
	case "CustomerService":
		return app.CustomerService
	case "CustomerServiceSession":
		return app.CustomerServiceSession
	case "Semantic":
		return app.Semantic
	case "DataCube":
		return app.DataCube
	case "AutoReply":
		return app.AutoReply
	case "Broadcasting":
		return app.Broadcasting
	case "Card":
		return app.Card
	case "Device":
		return app.Device
	case "ShakeAround":
		return app.ShakeAround
	case "POI":
		return app.POI
	case "Store":
		return app.Store
	case "Comment":
		return app.Comment
	case "OCR":
		return app.OCR
	case "Goods":
		return app.Goods
	case "OAuth":
		return app.OAuth
	case "Wifi":
		return app.Wifi
	case "WifiCard":
		return app.WifiCard
	case "WifiDevice":
		return app.WifiDevice
	case "WifiShop":
		return app.WifiShop
	case "Guide":
		return app.Guide

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
