package auth

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *AccessToken) {

	return  NewClient(&app),NewAccessToken(&app)

}
