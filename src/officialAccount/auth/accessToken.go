package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func NewAccessToken(app kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)
	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		kernelToken,
	}

	// Override fields and functions
	cfg := app.GetConfig()
	useStableToken := cfg.GetBool("stable_token_mode", false)
	forceRefresh := cfg.GetBool("force_refresh", false)
	// baseUrl := cfg.GetString("http.base_uri", "https://api.weixin.qq.com")
	// 使用 http.base_uri 拼接url 多一个 '/'，导致 接口请求 404
	baseUrl := "https://api.weixin.qq.com"
	if useStableToken {
		token.EndpointToGetToken = baseUrl + "/cgi-bin/stable_token"
		token.StableTokenMode = true
		token.ForceRefresh = forceRefresh
	} else {
		token.EndpointToGetToken = baseUrl + "/cgi-bin/token"
	}
	token.OverrideGetCredentials()

	return token, nil
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (accessToken.App).GetContainer().GetConfig()
	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"grant_type": "client_credential",
			"appid":      (*config)["app_id"].(string),
			"secret":     (*config)["secret"].(string),
			"neededText": "",
		}
	}
}
