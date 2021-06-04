package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
)


type AccessToken struct {
	App             *kernel.ServiceContainer
	BaseAccessToken *kernel.AccessToken
}

func Inject(app contract.ApplicationInterface){

	component := &AccessToken{
		App:             app.GetContainer(),
		BaseAccessToken: &kernel.AccessToken{},
	}
	components := app.GetComponents()

	(*components)["access_token"] = component

}

func (token *AccessToken) GetCredentials() object.StringMap {

	config := token.App.GetConfig()
	return object.StringMap{
		"corpid":     config[2]["corp_id"].(string),
		"corpsecret": config[2]["secret"].(string),
	}
}
