package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface) *AccessToken {
	token := &AccessToken{
		kernel.NewAccessToken(app),
	}

	// Override fields and functions
	token.EndpointToGetToken = "https://api.weixin.qq.com/cgi-bin/token"
	token.OverrideGetCredentials()

	return token
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetContainer().GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"grant_type": "client_credential",
			"appid":      (*config)["app_id"].(string),
			"secret":     (*config)["secret"].(string),
		}
	}
}
