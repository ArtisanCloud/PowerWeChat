package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *AccessToken {

	return NewAccessToken(&app)

}

func RegisterAuthProvider(app kernel.ApplicationInterface) *Client {

	return NewClient(&app)
}
