package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *AccessToken, error) {
	clt, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}
	accessToken, err := NewAccessToken(app)
	if err != nil {
		return nil, nil, err
	}
	return clt, accessToken, nil

}
