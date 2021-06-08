package auth

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func Inject(app kernel.ApplicationInterface) {
	container := app.GetContainer()
	container.GetConfig()

	component := NewAccessToken(&app)
	components := app.GetComponents()

	(*components)["access_token"] = component

}
