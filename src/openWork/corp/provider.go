package corp

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface, corpID string, permanentCode string) (*AccessToken, error) {
	return NewAccessToken(&app, corpID, permanentCode)
}
