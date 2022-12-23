package groupRobot

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Messager, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	messager := NewMessager(client)

	return client, messager, nil

}
