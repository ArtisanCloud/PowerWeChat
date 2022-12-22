package accountService

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/customer"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/message"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/servicer"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*customer.Client,
	*message.Client,
	*servicer.Client,
	*serviceState.Client,
	*tag.Client,
	error,
) {
	//config := app.GetConfig()

	Client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	Customer, err := customer.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	Message, err := message.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	Servicer, err := servicer.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	ServiceState, err := serviceState.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	Tag, err := tag.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return Client,
		Customer,
		Message,
		Servicer,
		ServiceState,
		Tag,
		nil
}
