package auth

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *AccessToken {

	return NewAccessToken(&app)

}
