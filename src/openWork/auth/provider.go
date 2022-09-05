package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	accessToken, err := NewAccessToken(&app)

	return accessToken, err

}
