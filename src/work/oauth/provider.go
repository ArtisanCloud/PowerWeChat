package oauth

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/exception"
	"strings"
)

func RegisterProvider(app kernel.ApplicationInterface) *Manager {

	config := *app.GetContainer().Config

	socialite := NewManager(&object.HashMap{
		"wecom": &object.HashMap{
			"client_id":     config["corp_id"].(string),
			"client_secret": "",
			"corp_id":       config["corp_id"].(string),
			"corp_secret":   config["secret"].(string),
			"redirect":      prepareCallbackUrl(app),
		},
	}, &app)

	globalConfig := app.GetConfig()
	scopes := globalConfig.Get("oauth.scopes", []string{"snsapi_base"}).([]string)

	if len(scopes) > 0 {
		socialite.Provider.Scopes(scopes)
	} else {
		agentID := globalConfig.Get("agent_id", 0).(int)
		socialite.Provider.SetAgentID(agentID)
	}

	return socialite

}

func prepareCallbackUrl(app kernel.ApplicationInterface) string {
	config := *app.GetContainer().Config

	callback := config["oauth.callback"].(string)
	if strings.Index(callback, "http") == 0 {
		return callback
	} else {
		// have to setup a complete url with host
		defer (&exception.Exception{}).HandleException(nil, "oauth.prepare.callback.url", nil)
		panic(fmt.Sprintf("OAuth callback format invalid, please make sure that schema 'http' added: %v", callback))

	}

}
