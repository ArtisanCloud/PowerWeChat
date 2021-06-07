package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

func NewAccessToken(container *kernel.ServiceContainer) *AccessToken {
	token := &AccessToken{
		kernel.NewAccessToken(container),
	}

	// Override fields and functions
	token.EndpointToGetToken = "cgi-bin/gettoken"
	token.OverrideGetCredentials()

	return token
}


// Override GetCredentials
func (component *AccessToken) OverrideGetCredentials() {
	config := component.App.GetConfig()
	component.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"corpid":     config[kernel.CONFIG_USER_INDEX]["corp_id"].(string),
			"corpsecret": config[kernel.CONFIG_USER_INDEX]["secret"].(string),
		}
	}
}
