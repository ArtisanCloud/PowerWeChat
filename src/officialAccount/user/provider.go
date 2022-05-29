package user

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/user/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *tag.Client) {

	userClient := NewClient(app)

	tagClient := tag.NewClient(app)

	return userClient, tagClient
}
