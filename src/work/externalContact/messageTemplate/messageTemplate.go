package messageTemplate

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type MessageTemplate struct {
	*kernel.BaseClient
}

func NewMessageTemplate(app kernel.ApplicationInterface) *MessageTemplate {
	return &MessageTemplate{
		kernel.NewBaseClient(&app, nil),
	}
}
