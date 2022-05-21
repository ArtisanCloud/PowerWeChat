package shakeAround

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type PageClient struct {
	*kernel.BaseClient
}


func NewPageClient(app kernel.ApplicationInterface) *PageClient {
	return &PageClient{
		kernel.NewBaseClient(&app, nil),
	}
}