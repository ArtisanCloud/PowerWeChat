package user

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	clt, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return clt, nil
}
