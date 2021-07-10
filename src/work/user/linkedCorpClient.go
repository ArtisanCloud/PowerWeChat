package user

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type LinkedCorpClient struct {
	*kernel.BaseClient
}

func NewLinkedCorpClient(app kernel.ApplicationInterface) *LinkedCorpClient {
	return &LinkedCorpClient{
		kernel.NewBaseClient(&app, nil),
	}
}
