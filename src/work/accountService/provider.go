package accountService

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/accountService/customer"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/accountService/message"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/accountService/servicer"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/accountService/tag"
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
