package auth

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
)

func Inject(app contract.ApplicationInterface) {
	container := app.GetContainer()
	container.GetConfig()

	component := NewAccessToken(container)
	components := app.GetComponents()

	(*components)["access_token"] = component

}
