package wxaCode

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(&app, nil),
	}

}
