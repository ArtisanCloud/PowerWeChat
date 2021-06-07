package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type AccessToken struct {
	BaseAccessToken *kernel.AccessToken

}

func NewAccessToken(container *kernel.ServiceContainer) *AccessToken {
	token := &AccessToken{
		BaseAccessToken: kernel.NewAccessToken(container),
	}

	token.BaseAccessToken.EndpointToGetToken = "cgi-bin/gettoken"
	token.SetupCredentials()

	return token
}



func (component *AccessToken) SetupCredentials(){
	config := component.BaseAccessToken.App.GetConfig()
	component.BaseAccessToken.GetCredentials = func() *object.StringMap {

		return &object.StringMap{
			"corpid":     config[kernel.CONFIG_USER_INDEX]["corp_id"].(string),
			"corpsecret": config[kernel.CONFIG_USER_INDEX]["secret"].(string),
		}
	}

}
