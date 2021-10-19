package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

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
	config := (*accessToken.App).GetConfig()

	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"grant_type": "client_credential",
			"appid":      config.GetString("app_id", ""),
			"secret":     config.GetString("secret", ""),
		}
	}
}
