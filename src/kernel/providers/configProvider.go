package providers

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterConfigProvider(app kernel.ApplicationInterface) *kernel.Config {

	return kernel.NewConfig(app.GetContainer().GetConfig())
}
