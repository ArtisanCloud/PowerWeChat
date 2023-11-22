package corp

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suit "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)

	kernelToken.RequestMethod = http.MethodPost
	kernelToken.TokenKey = "access_token"
	kernelToken.EndpointToGetToken = "cgi-bin/service/get_corp_token"
	kernelToken.CachePrefix = "powerwechat.kernel.get_corp_token."

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

	app := (*accessToken.App)
	accessToken.GetCredentials = func() *object.StringMap {
		config := app.GetContainer().GetConfig()
		token := app.GetComponent("SuiteAccessToken").(*suit.AccessToken)
		suiteAccessToken, _ := token.AccessToken.GetToken(false)

		corpID := (*config)["corp_id"].(string)
		appID := (*config)["app_id"].(string)
		permanentCode := (*config)["permanent_code"].(string)

		return &object.StringMap{
			"auth_corpid":        corpID,
			"permanent_code":     permanentCode,
			"suite_access_token": suiteAccessToken.AccessToken,

			"appid":      corpID,
			"secret":     permanentCode,
			"neededText": security.HashStringData(corpID + "-" + appID),
		}
	}
}
