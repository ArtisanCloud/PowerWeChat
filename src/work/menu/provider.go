package menu

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	client, err := NewClient(app)

	return client, err

}
