package suit

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*SuiteTicket, *AccessToken, error) {

	ticket, err := NewSuiteTicket(&app)
	accessToken, err := NewAccessToken(&app)

	return ticket, accessToken, err

}
