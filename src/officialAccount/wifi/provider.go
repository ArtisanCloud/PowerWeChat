package wifi

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *CardClient, *DeviceClient, *ShopClient) {

	client := NewClient(app)
	cardClient := NewCardClient(app)
	deviceClient := NewDeviceClient(app)
	shopClient := NewShopClient(app)

	return client, cardClient, deviceClient, shopClient
}
