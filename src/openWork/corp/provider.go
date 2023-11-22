package corp

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	accessToken, err := NewAccessToken(&app)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
