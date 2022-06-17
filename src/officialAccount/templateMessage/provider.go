package templateMessage

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
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

	return client, nil
}
