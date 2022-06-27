package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91039
func NewAccessToken(app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)
	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		kernelToken,
	}

	// Override fields and functions
	token.EndpointToGetToken = "cgi-bin/gettoken"
	token.OverrideGetCredentials()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetContainer().GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			// this is for the cached key encoded
			"appid":      (*config)["corp_id"].(string),
			"secret":     (*config)["secret"].(string),
			"neededText": "",

			// this is for the real credentials
			"corpid":     (*config)["corp_id"].(string),
			"corpsecret": (*config)["secret"].(string),
		}
	}
}
