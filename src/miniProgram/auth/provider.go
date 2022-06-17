package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	return NewAccessToken(&app)

}

func RegisterAuthProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)
}
