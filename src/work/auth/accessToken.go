package auth

// #reference: https://developer.work.weixin.qq.com/document/path/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developer.work.weixin.qq.com/document/path/91039
func NewAccessToken(app kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(&app)
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
	config := (*accessToken.App).GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			// this is for the cached key encoded
			"appid":      config.GetString("corp_id", ""),
			"secret":     config.GetString("secret", ""),
			"neededText": "",

			// this is for the real credentials
			"corpid":     config.GetString("corp_id", ""),
			"corpsecret": config.GetString("secret", ""),
		}
	}
}
