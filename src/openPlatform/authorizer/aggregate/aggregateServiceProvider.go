package aggregate

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/aggregate/account"
)

func RegisterProvider(app kernel.ApplicationInterface) (*account.Client, error) {

	account, err := account.NewClient(app)
	if err != nil {
		return nil, err
	}

	return account, nil
}
