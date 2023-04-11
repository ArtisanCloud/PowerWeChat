package payment

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/bill"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/jssdk"
	kernel2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/partner"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/redpack"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/reverse"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/sandbox"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/security"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer"
	"net/http"
	"time"
)

type Payment struct {
	kernel2.ApplicationPaymentInterface
	*kernel.ServiceContainer

	Config *kernel.Config

	Order   *order.Client
	Partner *partner.Client
	JSSDK   *jssdk.Client
	Sandbox *sandbox.Client

	Refund *refund.Client
	Bill   *bill.Client

	RedPack       *redpack.Client
	Transfer      *transfer.Client
	TransferBatch *transfer.BatchClient
	Reverse       *reverse.Client
	ProfitSharing *profitSharing.Client

	Base *base.Client

	Security *security.Client

	Logger *logger.Logger
}

type UserConfig struct {
	AppID              string
	MchID              string
	MchApiV3Key        string
	Key                string
	CertPath           string
	KeyPath            string
	SerialNo           string
	CertificateKeyPath string
	WechatPaySerial    string
	RSAPublicKeyPath   string
	SpAppID            string
	SpMchID            string
	SubAppID           string
	SubMchID           string
	Http               Http

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
			"http": &object.HashMap{
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

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat.log"),
		"errorPath":  app.Config.GetString("log.file", "./wechat.log"),
	})

	if err != nil {
		return nil, err
	}
	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- Order --------------
	app.Order, err = order.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Partner --------------
	app.Partner, err = partner.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- JSSDK --------------
	app.JSSDK, err = jssdk.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Sandbox --------------
	app.Sandbox, err = sandbox.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Refund --------------
	app.Refund, err = refund.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Bill --------------
	app.Bill, err = bill.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Red Pack --------------
	app.RedPack, err = redpack.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- Transfer --------------
	app.Transfer, app.TransferBatch, err = transfer.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- Reverse --------------
	app.Reverse, err = reverse.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- ProfitSharing --------------
	app.ProfitSharing, err = profitSharing.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- Security --------------
	app.Security, err = security.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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
	case "Partner":
		return app.Partner
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

	case "Security":
		return app.Security

	case "Logger":
		return app.Logger
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

func (app *Payment) SetSubMerchant(mchID string, appID string) kernel2.ApplicationPaymentInterface {
	app.Config.Set("sub_mch_id", mchID)
	app.Config.Set("sub_appid", appID)

	return app
}

func (app *Payment) HandlePaidNotify(request *http.Request, closure func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{}) (*http.Response, error) {
	return notify.NewPaidNotify(app, request).Handle(closure)
}

func (app *Payment) HandleRefundedNotify(request *http.Request, closure func(message *request.RequestNotify, transaction *models.Refund, fail func(message string)) interface{}) (*http.Response, error) {
	return notify.NewRefundNotify(app, request).Handle(closure)
}

func (app *Payment) HandleScannedNotify(request *http.Request, closure func(message *request.RequestNotify, fail func(message string), alert func(message string)) interface{}) (*http.Response, error) {
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
		"app_id":               userConfig.AppID,
		"mch_id":               userConfig.MchID,
		"mch_api_v3_key":       userConfig.MchApiV3Key,
		"key":                  userConfig.Key,
		"cert_path":            userConfig.CertPath,
		"key_path":             userConfig.KeyPath,
		"certificate_key_path": userConfig.CertificateKeyPath,
		"serial_no":            userConfig.SerialNo,
		"rsa_public_key_path":  userConfig.RSAPublicKeyPath,
		"wechat_pay_serial":    userConfig.WechatPaySerial,
		"sub_appid":            userConfig.SubAppID,
		"sub_mch_id":           userConfig.SubMchID,

		"response_type": userConfig.ResponseType,
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
		},
		"http": &object.HashMap{
			"timeout":  userConfig.Http.Timeout,
			"base_uri": userConfig.Http.BaseURI,
		},
		"oauth.callbacks": userConfig.OAuth.Callback,
		"oauth.scopes":    userConfig.OAuth.Scopes,
		"notify_url":      userConfig.NotifyURL,
		"cache":           userConfig.Cache,
		"http_debug":      userConfig.HttpDebug,
		"debug":           userConfig.Debug,
		"sandbox":         userConfig.Sandbox,
	}

	return config, nil

}
