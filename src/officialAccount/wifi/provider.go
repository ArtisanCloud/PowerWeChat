package wifi

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *CardClient, *DeviceClient, *ShopClient) {

	client := NewClient(app)
	cardClient := NewCardClient(app)
	deviceClient := NewDeviceClient(app)
	shopClient := NewShopClient(app)

	return client, cardClient, deviceClient, shopClient
}
