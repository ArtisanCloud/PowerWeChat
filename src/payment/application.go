package payment

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	"github.com/ArtisanCloud/power-wechat/src/payment/notify"
	"github.com/ArtisanCloud/power-wechat/src/payment/sandbox"
	"github.com/ArtisanCloud/power-wechat/src/payment/base"
	"net/http"
	"time"
)

type Payment struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request
	Config          *kernel.Config

	Sandbox *sandbox.Client

	Base *base.Client
}

func NewPayment(config *object.HashMap, r *http.Request) (*Payment, error) {
	var err error

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: config,
		DefaultConfig: &object.HashMap{
			"http": map[string]string{
				"base_uri": "https://api.mch.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Payment{
		ServiceContainer: container,
	}

	//-------------- external request --------------
	app.ExternalRequest = r
	if r == nil {
		app.ExternalRequest = &http.Request{}
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

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
	case "ExternalRequest":
		return app.ExternalRequest
	case "Base":
		return app.Base
	//case "AccessToken":
	//	return app.AccessToken
	case "Config":
		return app.Config

	default:
		return nil
	}

}

func (app *Payment) Scheme(productID string) string {
	appID := app.Config.GetString("app_id", "")
	mchID := app.Config.GetString("mch_id", "")
	key := app.Config.GetString("key", "")
	params := &object.StringMap{
		"appid":      appID,
		"mch_id":     mchID,
		"time_stamp": fmt.Sprintf("%d", time.Now().Nanosecond()),
		"nonce_str":  str.UniqueID(""),
		"product_id": productID,
	}

	(*params)["sign"] = support.GenerateSign(params, key, "md5")

	return "weixin://wxpay/bizpayurl?" + object.ConvertStringMapToString(params, "&")
}

func (app *Payment) CodeUrlScheme(codeUrl string) string {
	return fmt.Sprintf("weixin://wxpay/bizpayurl?sr=%s", codeUrl)
}

func (app *Payment) SetSubMerchant(mchId string, appId string) *Payment {
	app.Config.Set("sub_mch_id", mchId)
	app.Config.Set("sub_appid", appId)

	return app
}

func (app *Payment) HandlePaidNotify(closure func(payload ...interface{}) interface{}) {
	notify.NewPaid(app).Handle(closure)
}

func (app *Payment) HandleRefundedNotify(closure func(payload ...interface{}) interface{}) {
	notify.NewRefund(app).Handle(closure)
}

func (app *Payment) HandleScannedNotify(closure func(payload ...interface{}) interface{}) {
	notify.NewScanned(app).Handle(closure)
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
