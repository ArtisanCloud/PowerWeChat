package oauth

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"strings"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Manager, error) {

	config := *app.GetContainer().Config
	globalConfig := app.GetConfig()
	prepareCallbackURL, err := prepareCallbackUrl(app)
	if err != nil {
		return nil, err
	}
	managerConfig := &object.HashMap{
		"client_id":     config["corp_id"].(string),
		"client_secret": "",
		"corp_id":       config["corp_id"].(string),
		"corp_secret":   config["secret"].(string),
		"redirect":      prepareCallbackURL,
	}
	providerConfig := object.MergeHashMap(globalConfig.All(), managerConfig)
	socialite := NewManager(
		&object.HashMap{
			"wecom": managerConfig,
		}, providerConfig, &app)

	scopes := globalConfig.Get("oauth.scopes", []string{"snsapi_base"}).([]string)

	if len(scopes) > 0 {
		socialite.Provider.Scopes(scopes)
	} else {
		agentID := globalConfig.Get("agent_id", 0).(int)
		socialite.Provider.SetAgentID(agentID)
	}

	return socialite, nil

}

func prepareCallbackUrl(app kernel.ApplicationInterface) (string, error) {
	config := *app.GetContainer().Config

	var callback string
	if config["oauth.callback"] != nil {
		callback = config["oauth.callback"].(string)
		if strings.Index(callback, "http") == 0 {
			return callback, nil
		}
	}

	return callback, errors.New(fmt.Sprintf("OAuth callback format invalid, please make sure that schema 'http' added: %v", callback))

}
