package auth

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *AccessToken {

	return NewAccessToken(&app)

}
