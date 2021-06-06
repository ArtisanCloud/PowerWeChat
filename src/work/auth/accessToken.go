package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)


type AccessToken struct {
	App             *kernel.ServiceContainer
	BaseAccessToken *kernel.AccessToken
}



func (token *AccessToken) GetCredentials() object.StringMap {

	config := token.App.GetConfig()
	return object.StringMap{
		"corpid":     config[2]["corp_id"].(string),
		"corpsecret": config[2]["secret"].(string),
	}
}
