package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/go-wechat/src/kernel/http"
)

func Inject(app contract.ApplicationInterface) {
	container := app.GetContainer()
	container.GetConfig()

	component := &Client{
		BaseClient: &http.BaseClient{
			App: container,
		},
	}
	components := app.GetComponents()

	(*components)["base"] = component
}
