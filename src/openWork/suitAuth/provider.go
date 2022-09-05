package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*SuiteTicket, *AccessToken, error) {

	ticket, err := NewSuiteTicket(&app)
	accessToken, err := NewAccessToken(&app)

	return ticket, accessToken, err

}
