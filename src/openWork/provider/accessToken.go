package provider

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
	CorpID string `json:"corpid,omitempty"`
}

// https://developer.work.weixin.qq.com/document/path/90593
func NewAccessToken(app *kernel.ApplicationInterface, corpID string) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = http.MethodPost
	kernelToken.TokenKey = "provider_access_token"
	kernelToken.EndpointToGetToken = "cgi-bin/service/get_provider_token"
	kernelToken.CachePrefix = "powerwechat.kernel.provider_access_token."

	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		kernelToken,
		corpID,
	}

	// Override fields and functions
	token.OverrideGetCredentials()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {

	accessToken.GetCredentials = func() *object.StringMap {
		config := (*accessToken.App).GetContainer().GetConfig()

		corpID := (*config)["provider_corpid"].(string)
		secret := (*config)["provider_secret"].(string)

		return &object.StringMap{
			"corpid":          corpID,
			"provider_secret": secret,

			"appid":      corpID,
			"secret":     secret,
			"neededText": security.HashStringData(corpID + "-" + secret),
		}
	}
}
