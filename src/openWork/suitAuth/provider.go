package suit

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *SuiteTicket, *AccessToken, error) {

	clt, err := NewClient(&app)
	if err != nil {
		return nil, nil, nil, err
	}
	ticket, err := NewSuiteTicket(&app)
	if err != nil {
		return nil, nil, nil, err
	}
	accessToken, err := NewAccessToken(&app)
	if err != nil {
		return nil, nil, nil, err
	}

	return clt, ticket, accessToken, nil

}
