package templateMessage

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) *Client {

	client := &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),
		Message: &object.HashMap{
			"touser":      "",
			"template_id": "",
			"url":         "",
			"data":        &object.HashMap{},
			"miniprogram": "",
		},
		Required: []string{
			"touser",
			"template_id",
		},
	}

	return client
}
