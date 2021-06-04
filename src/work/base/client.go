package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func Inject(app contract.ApplicationInterface){

	component := &Client{
		BaseClient:             &kernel.BaseClient{},
	}
	components := app.GetComponents()

	(*components)["base"] = component

}