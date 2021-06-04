package auth

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type AccessToken struct {
	BaseAccessToken kernel.AccessToken
}

func (token *AccessToken)GetCredentials() object.StringMap {
	return object.StringMap{
		"corpid": ,
		"corpsecret": ,
	}
}