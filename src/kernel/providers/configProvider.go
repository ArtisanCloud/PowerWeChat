package providers

import "github.com/ArtisanCloud/go-wechat/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) *kernel.Config {

	return kernel.NewConfig(app.GetContainer().GetConfig())
}
