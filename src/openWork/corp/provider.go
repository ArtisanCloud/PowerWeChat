package corp

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	code, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return code, nil
}
