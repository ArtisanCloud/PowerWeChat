package work

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"net/http"
	"time"
)

type Payment struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request
	Config          *kernel.Config

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
	appID := app.Config.Get("app_id", "").(string)
	mchID := app.Config.Get("mch_id", "").(string)
	key := app.Config.Get("key", "").(string)
	params := &object.StringMap{
		"appid":      appID,
		"mch_id":     mchID,
		"time_stamp": fmt.Sprintf("%d", time.Now().Nanosecond()),
		"nonce_str":  str.UniqueID(""),
		"product_id": productID,
	}

	(*params)["sign"] = support.GenerateSign(params, key)

	return "weixin://wxpay/bizpayurl?" + object.ConvertStringMapToString(params, "&")
}
