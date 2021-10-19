package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91039
func NewAccessToken(app *kernel.ApplicationInterface) *AccessToken {
	token := &AccessToken{
		kernel.NewAccessToken(app),
	}

	// Override fields and functions
	token.EndpointToGetToken = "cgi-bin/gettoken"
	token.OverrideGetCredentials()

	return token
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetContainer().GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"corpid":     (*config)["corp_id"].(string),
			"corpsecret": (*config)["secret"].(string),
		}
	}
}
