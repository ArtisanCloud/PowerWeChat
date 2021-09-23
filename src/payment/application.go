package payment

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/payment/base"
	"github.com/ArtisanCloud/power-wechat/src/payment/bill"
	"github.com/ArtisanCloud/power-wechat/src/payment/jssdk"
	kernel2 "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/notify"
	"github.com/ArtisanCloud/power-wechat/src/payment/notify/request"
	"github.com/ArtisanCloud/power-wechat/src/payment/order"
	"github.com/ArtisanCloud/power-wechat/src/payment/profitSharing"
	"github.com/ArtisanCloud/power-wechat/src/payment/redpack"
	"github.com/ArtisanCloud/power-wechat/src/payment/refund"
	"github.com/ArtisanCloud/power-wechat/src/payment/reverse"
	"github.com/ArtisanCloud/power-wechat/src/payment/sandbox"
	"github.com/ArtisanCloud/power-wechat/src/payment/transfer"
	"net/http"
	"time"
)

type Payment struct {
	kernel2.ApplicationPaymentInterface
	*kernel.ServiceContainer

	Config *kernel.Config

	Order   *order.Client
	JSSDK   *jssdk.Client
	Sandbox *sandbox.Client

	Refund *refund.Client
	Bill   *bill.Client

	RedPack       *redpack.Client
	Transfer      *transfer.Client
	Reverse       *reverse.Client
	ProfitSharing *profitSharing.Client

	Base *base.Client
}

type UserConfig struct {
	AppID       string
	MchID       string
	MchApiV3Key string
	Key         string
	CertPath    string
	KeyPath     string
	SerialNo    string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Http         Http
	NotifyURL    string
	HttpDebug    bool
	Debug        bool
	Sandbox      bool
}

type Log struct {
	Level string
	File  string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

type Http struct {
	Timeout float64
	BaseURI string
}

func NewPayment(config *UserConfig) (*Payment, error) {
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
				"base_uri": "https://api.mch.weixin.qq.com",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Payment{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- Order --------------
	app.Order = order.RegisterProvider(app)

	//-------------- JSSDK --------------
	app.JSSDK = jssdk.RegisterProvider(app)

	//-------------- Sandbox --------------
	app.Sandbox = sandbox.RegisterProvider(app)

	//-------------- Refund --------------
	app.Refund = refund.RegisterProvider(app)

	//-------------- Bill --------------
	app.Bill = bill.RegisterProvider(app)

	//-------------- Red Pack --------------
	app.RedPack = redpack.RegisterProvider(app)

	//-------------- Transfer --------------
	app.Transfer = transfer.RegisterProvider(app)

	//-------------- Reverse --------------
	app.Reverse = reverse.RegisterProvider(app)

	//-------------- Reverse --------------
	app.ProfitSharing = profitSharing.RegisterProvider(app)

	return app, err
}

func (app *Payment) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Payment) GetAccessToken() *kernel.AccessToken {
	//return app.AccessToken.AccessToken
	return nil
}

func (app *Payment) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Payment) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "JSSDK":
		return app.JSSDK
	case "Sandbox":
		return app.Sandbox
	case "Config":
		return app.Config
	case "Order":
		return app.Order
	case "Refund":
		return app.Refund
	case "Bill":
		return app.Bill
	case "RedPack":
		return app.RedPack
	case "Transfer":
		return app.Transfer
	case "Reverse":
		return app.Reverse
	case "ProfitSharing":
		return app.ProfitSharing
	default:
		return nil
	}

}

func (app *Payment) Scheme(productID string) string {
	appID := app.Config.GetString("app_id", "")
	mchID := app.Config.GetString("mch_id", "")
	//key := app.Config.GetString("key", "")
	params := &object.StringMap{
		"appid":      appID,
		"mch_id":     mchID,
		"time_stamp": fmt.Sprintf("%d", time.Now().Nanosecond()),
		"nonce_str":  object.UniqueID(""),
		"product_id": productID,
	}

	//var err error
	//(*params)["sign"], err = support.GenerateSign(params, key, "md5")

	return "weixin://wxpay/bizpayurl?" + object.ConvertStringMapToString(params, "&")
}

func (app *Payment) CodeUrlScheme(codeUrl string) string {
	return fmt.Sprintf("weixin://wxpay/bizpayurl?sr=%s", codeUrl)
}

func (app *Payment) SetSubMerchant(mchId string, appId string) kernel2.ApplicationPaymentInterface {
	app.Config.Set("sub_mch_id", mchId)
	app.Config.Set("sub_appid", appId)

	return app
}

func (app *Payment) HandlePaidNotify(request *http.Request, closure func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{}) (*response.HttpResponse, error) {
	return notify.NewPaidNotify(app, request).Handle(closure)
}

func (app *Payment) HandleRefundedNotify(request *http.Request, closure func(message *request.RequestNotify, transaction *models.Refund, fail func(message string)) interface{}) (*response.HttpResponse, error) {
	return notify.NewRefundNotify(app, request).Handle(closure)
}

func (app *Payment) HandleScannedNotify(request *http.Request, closure func(message *request.RequestNotify, fail func(message string), alert func(message string)) interface{}) (*response.HttpResponse, error) {
	return notify.NewScannedNotify(app, request).Handle(closure)
}

func (app *Payment) InSandbox() bool {
	return app.Config.GetBool("sandbox", false)

}

func (app *Payment) GetKey(endpoint string) (string, error) {
	if "sandboxnew/pay/getsignkey" == endpoint {
		return app.Config.GetString("key", ""), nil
	}

	key := app.Config.GetString("key", "")
	if app.InSandbox() {
		key, _ = app.Sandbox.GetKey()
	}
	if key == "" {
		return key, errors.New("config key should not empty. ")
	}

	if len(key) != 32 {
		return key, errors.New(fmt.Sprintf("'%s' should be 32 chars length.", key))
	}

	return key, nil

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"app_id":         userConfig.AppID,
		"mch_id":         userConfig.MchID,
		"mch_api_v3_key": userConfig.MchApiV3Key,
		"key":            userConfig.Key,
		"cert_path":      userConfig.CertPath,
		"key_path":       userConfig.KeyPath,
		"serial_no":      userConfig.SerialNo,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"http": object.HashMap{
			"timeout":  userConfig.Http.Timeout,
			"base_uri": userConfig.Http.BaseURI,
		},
		"oauth.callback": userConfig.OAuth.Callback,
		"oauth.scopes":   userConfig.OAuth.Scopes,
		"notify_url":     userConfig.NotifyURL,
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
		"sandbox":        userConfig.Sandbox,
	}

	return config, nil

}
