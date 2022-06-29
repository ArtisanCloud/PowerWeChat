package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/auth"
)

type AccessToken struct {
	*kernel.AccessToken

	// PowerWechat\OpenPlatform\Application
	component kernel.ApplicationInterface
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface, component kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = "POST"
	kernelToken.QueryName = "access_token"
	kernelToken.TokenKey = "authorizer_access_token"

	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		AccessToken: kernelToken,
	}

	token.component = component

	// Override fields and functions
	token.OverrideGetCredentials()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetContainer().GetConfig()
	componentConfig := (*accessToken.App).GetContainer().GetConfig()

	accessToken.GetCredentials = func() *object.StringMap {

		return &object.StringMap{
			"component_appid":          (*componentConfig)["app_id"].(string),
			"authorizer_appid":         (*config)["app_id"].(string),
			"authorizer_refresh_token": (*config)["refresh_token"].(string),

			"appid":      (*componentConfig)["app_id"].(string),
			"secret":     (*config)["refresh_token"].(string),
			"neededText": (*config)["app_id"].(string),
		}
	}
}

// Override GetEndpoint
func (accessToken *AccessToken) OverrideGetEndpoint() {
	accessToken.GetEndpoint = func() (string, error) {
		token := accessToken.component.GetComponent("AccessToken").(*auth.AccessToken)
		componentToken, err := token.GetToken(false)
		if err != nil {
			return "", err
		}
		return "cgi-bin/component/api_authorizer_token?" + object.GetJoinedWithKSort(&object.StringMap{
			"component_access_token": componentToken.ComponentAccessToken,
		}), nil
	}
}
