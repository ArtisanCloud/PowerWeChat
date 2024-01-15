package corp

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"net/http"
	"net/url"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suit "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
)

type AccessToken struct {
	*kernel.AccessToken
	corpID        string
	permanentCode string
}

// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func NewAccessToken(app *kernel.ApplicationInterface, corpId, permanentCode string) (*AccessToken, error) {
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
		corpId,
		permanentCode,
	}

	// Override fields and functions
	token.OverrideGetCredentials()
	token.OverrideGetEndpoint()

	return token, nil
}

func (accessToken *AccessToken) OverrideGetEndpoint() {

	app := (*accessToken.App)
	accessToken.GetEndpoint = func() (string, error) {
		token := app.GetComponent("SuiteAccessToken").(*suit.AccessToken)
		suiteAccessToken, _ := token.AccessToken.GetToken(false)
		return accessToken.EndpointToGetToken + `?suite_access_token=` + url.QueryEscape(suiteAccessToken.AccessToken), nil
	}

}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {

	app := (*accessToken.App)
	accessToken.GetCredentials = func() *object.StringMap {
		config := app.GetContainer().GetConfig()

		appID := (*config)["app_id"].(string)

		return &object.StringMap{
			"auth_corpid":    accessToken.corpID,
			"permanent_code": accessToken.permanentCode,

			"appid":      accessToken.corpID,
			"secret":     accessToken.permanentCode,
			"neededText": security.HashStringData(accessToken.corpID + "-" + appID),
		}
	}
}
