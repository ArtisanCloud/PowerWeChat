package auth

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*VerifyTicket, *AccessToken, error) {

	ticket, err := NewVerifyTicket(&app)
	accessToken, err := NewAccessToken(app)

	return ticket, accessToken, err

}
