package accountService

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/customer"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/message"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/servicer"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*customer.Client,
	*message.Client,
	*servicer.Client,
	*serviceState.Client,
	*tag.Client,
) {
	//config := app.GetConfig()

	Client := NewClient(app)
	Customer := customer.NewClient(app)
	Message := message.NewClient(app)
	Servicer := servicer.NewClient(app)
	ServiceState := serviceState.NewClient(app)
	Tag := tag.NewClient(app)

	return Client,
		Customer,
		Message,
		Servicer,
		ServiceState,
		Tag
}
