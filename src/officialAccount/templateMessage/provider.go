package templateMessage

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
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
