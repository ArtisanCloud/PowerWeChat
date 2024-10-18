package provider

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {
	//client, err := NewClient(app)
	//if err != nil {
	//	return nil, nil, err
	//}
	accessToken, err := NewAccessToken(app)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
