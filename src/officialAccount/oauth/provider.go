package oauth

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerSocialite/v2/src"
	"github.com/ArtisanCloud/PowerSocialite/v2/src/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"strings"
)

func RegisterProvider(app kernel.ApplicationInterface) *providers.WeChat {

	config := *app.GetConfig()

	wechat := &object.HashMap{
		"wechat": &object.HashMap{
			"client_id":     config.GetString("app_id", ""),
			"client_secret": config.GetString("secret", ""),
			"redirect":      prepareCallbackURL(app),
			"http_debug":    config.GetBool("http_debug", false),
			"debug":         config.GetBool("http_debug", false),
		},
	}

	if config.GetString("component_app_id", "") != "" &&
		config.GetString("component_app_token", "") != "" {

		(*(*wechat)["wechat"].(*object.HashMap))["component"] = &object.HashMap{
			"id":    config.GetString("component_app_id", ""),
			"token": config.GetString("token", ""),
		}
	}

	socialite := src.NewSocialiteManager(wechat)
	socialite.Create("wechat")
	managerProviders := socialite.GetResolvedProviders()
	wechatProvider := (*managerProviders)["wechat"].(*providers.WeChat)

	scopes := config.Get("oauth.scopes", []string{"snsapi_base"}).([]string)

	if len(scopes) > 0 {
		wechatProvider.Scopes(scopes)
	}

	return wechatProvider
}

func prepareCallbackURL(app kernel.ApplicationInterface) string {
	config := *app.GetConfig()

	var callback string
	callback = config.GetString("oauth.callbacks", "")
	if callback != "" {
		if strings.Index(callback, "http") == 0 {
			return callback
		}
	}

	//externalRequest := app.GetExternalRequest()
	baseURL := config.GetString("oauth.callbacks", "")
	//if externalRequest != nil {
	//	baseURL = app.GetExternalRequest().URL.Host
	//}

	return baseURL + "/" + callback

}
