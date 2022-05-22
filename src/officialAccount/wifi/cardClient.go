package wifi

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type CardClient struct {
	*kernel.BaseClient
}


func NewCardClient(app kernel.ApplicationInterface) *CardClient {
	return &CardClient{
		kernel.NewBaseClient(&app, nil),
	}
}