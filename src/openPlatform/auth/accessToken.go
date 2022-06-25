package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = "POST"
	kernelToken.TokenKey = "component_access_token"
	kernelToken.EndpointToGetToken = "cgi-bin/component/api_component_token"

	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		kernelToken,
	}

	// Override fields and functions
	token.OverrideGetCredentials()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetContainer().GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {

		ticket, _ := (*config)["verify_ticket"].(*VerifyTicket).GetTicket()

		return &object.StringMap{
			"component_appid":         (*config)["app_id"].(string),
			"component_appsecret":     (*config)["secret"].(string),
			"component_verify_ticket": ticket,

			"appid":  (*config)["corp_id"].(string),
			"secret": (*config)["secret"].(string),
			"ticket": ticket,
		}
	}
}
