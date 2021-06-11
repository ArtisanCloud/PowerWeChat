package auth

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *AccessToken {

	return NewAccessToken(&app)

}
