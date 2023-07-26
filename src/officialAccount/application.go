package officialAccount

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	providers2 "github.com/ArtisanCloud/PowerSocialite/v3/src/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/media"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/qrCode"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/url"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/auth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/autoReply"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/broadcasting"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/card"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/comment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/customerService"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/customerService/session"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/dataCube"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/device"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/goods"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/guide"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/material"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/menu"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/oauth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/ocr"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/poi"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/publish"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/semantic"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/store"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user/tag"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi"
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
	Publish                *publish.Client
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

	Logger *logger.Logger
}

type UserConfig struct {
	AppID        string
	Secret       string
	Token        string
	AESKey       string
	RefreshToken string

	ComponentAppID    string
	ComponentAppToken string

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
	Timeout float64
	BaseURI string
}

type Log struct {
	Level string
	File  string
	Error string
	ENV   string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewOfficialAccount(config *UserConfig, extraInfos ...*kernel.ExtraInfo) (*OfficialAccount, error) {
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
	app := &OfficialAccount{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat/info.log"),
		"errorPath":  app.Config.GetString("log.error", "./wechat/error.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- media --------------
	app.Media, err = media.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register QRCode --------------
	app.QRCode, err = qrCode.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register URL --------------
	app.URL, err = url.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register JSSDK --------------
	app.JSSDK, err = jssdk.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register SubscribeMessage --------------
	app.SubscribeMessage, err = subscribeMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Menu --------------
	app.Menu, err = menu.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Material --------------
	app.Material, err = material.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register CustomerService --------------
	app.CustomerService, app.CustomerServiceSession, err = customerService.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Semantic --------------
	app.Semantic, err = semantic.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Encryptor and Server --------------
	app.Encryptor, app.Server, err = server.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register User --------------
	app.User, app.UserTag, err = user.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Menu --------------
	app.Menu, err = menu.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register TemplateMessage --------------
	app.TemplateMessage, err = templateMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Material --------------
	app.Material, err = material.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register CustomerService --------------
	app.CustomerService, app.CustomerServiceSession, err = customerService.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Semantic --------------
	app.Semantic, err = semantic.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register DataCube --------------
	app.DataCube, err = dataCube.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register AutoReply --------------
	app.AutoReply, err = autoReply.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Broadcasting --------------
	app.Broadcasting, err = broadcasting.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Card --------------
	app.Card, err = card.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Device --------------
	app.Device, err = device.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register ShakeAround --------------
	app.ShakeAround, err = shakeAround.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register POI --------------
	app.POI, err = poi.RegisterProvider(app)
	//-------------- register Publish --------------
	app.Publish, err = publish.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Store --------------
	app.Store, err = store.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Comment --------------
	app.Comment, err = comment.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register OCR --------------
	app.OCR, err = ocr.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Goods --------------
	app.Goods, err = goods.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register OAuth --------------
	app.OAuth = oauth.RegisterProvider(app)
	//-------------- register wifi --------------
	app.Wifi, app.WifiCard, app.WifiDevice, app.WifiShop, err = wifi.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Guide --------------
	app.Guide, err = guide.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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
	case "Publish":
		return app.Publish
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

		"app_id":        userConfig.AppID,
		"secret":        userConfig.Secret,
		"token":         userConfig.Token,
		"aes_key":       userConfig.AESKey,
		"refresh_token": userConfig.RefreshToken,

		"component_app_id":    userConfig.ComponentAppID,
		"component_app_token": userConfig.ComponentAppToken,

		"response_type": userConfig.ResponseType,
		"http": &object.HashMap{
			"timeout":  timeout,
			"base_uri": baseURI,
		},
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
