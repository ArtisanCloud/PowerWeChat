package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/auth"
	"net/http"
)

type AccessToken struct {
	*kernel.AccessToken

	// PowerWechat\OpenPlatform\Application
	Component kernel.ApplicationInterface
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func NewAccessToken(app kernel.ApplicationInterface, component kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = http.MethodPost
	kernelToken.QueryName = "access_token"
	kernelToken.TokenKey = "authorizer_access_token"

	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		AccessToken: kernelToken,
	}

	token.Component = component

	// Override fields and functions
	token.OverrideGetCredentials()
	token.OverrideGetEndpoint()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {

	accessToken.GetCredentials = func() *object.StringMap {

		config := (accessToken.App).GetContainer().GetConfig()
		componentConfig := accessToken.Component.GetContainer().GetConfig()
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
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (accessToken *AccessToken) OverrideGetEndpoint() {
	accessToken.GetEndpoint = func() (string, error) {
		token := accessToken.Component.GetComponent("AccessToken").(*auth.AccessToken)
		componentToken, err := token.GetToken(context.Background(), false)
		if err != nil {
			return "", err
		}
		return "cgi-bin/component/api_authorizer_token?" + object.GetJoinedWithKSort(&object.StringMap{
			"component_access_token": componentToken.ComponentAccessToken,
		}), nil
	}
}
