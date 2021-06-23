package externalContact

import "github.com/ArtisanCloud/go-wechat/src/kernel"

type MessageTemplateClient struct {
	*kernel.BaseClient
}

func NewMessageTemplateClient(app kernel.ApplicationInterface) *MessageTemplateClient {
	return &MessageTemplateClient{
		kernel.NewBaseClient(&app, nil),
	}
}