package wifi

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *CardClient, *DeviceClient, *ShopClient, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	cardClient, err := NewCardClient(app)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	deviceClient, err := NewDeviceClient(app)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	shopClient, err := NewShopClient(app)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return client, cardClient, deviceClient, shopClient, nil
}
