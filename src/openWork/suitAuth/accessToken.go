package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = "POST"
	kernelToken.TokenKey = "suite_access_token"
	kernelToken.EndpointToGetToken = "cgi-bin/service/get_suite_token"
	kernelToken.CachePrefix = "powerwechat.kernel.suite_access_token."

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

	accessToken.GetCredentials = func() *object.StringMap {
		config := (*accessToken.App).GetContainer().GetConfig()
		suiteTicket := (*accessToken.App).GetComponent("SuiteTicket").(*SuiteTicket)
		ticket, _ := suiteTicket.GetTicket()

		return &object.StringMap{
			"suite_id":     (*config)["app_id"].(string),
			"suite_secret": (*config)["secret"].(string),
			"suite_ticket": ticket,

			"appid":      (*config)["app_id"].(string),
			"secret":     (*config)["secret"].(string),
			"neededText": ticket,
		}
	}
}
