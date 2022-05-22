package wifi

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type ShopClient struct {
	*kernel.BaseClient
}


func NewShopClient(app kernel.ApplicationInterface) *ShopClient {
	return &ShopClient{
		kernel.NewBaseClient(&app, nil),
	}
}