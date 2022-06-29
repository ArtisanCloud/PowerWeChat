package account

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	account, err := NewClient(app)
	if err != nil {
		return nil, err
	}

	return account, nil
}
