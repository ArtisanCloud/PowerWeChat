package license

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Account, error) {

	code, err := NewClient(&app)
	if err != nil {
		return nil, nil, err
	}

	account, err := NewAccount(&app)
	if err != nil {
		return nil, nil, err
	}

	return code, account, nil
}
