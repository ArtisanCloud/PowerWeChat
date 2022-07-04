package handlers

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/contract"
	"net/http"
)

type Authorized struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewAuthorized(app *kernel.ApplicationInterface) *Authorized {
	handler := &Authorized{
		App: app,
	}
	return handler
}

func (handler *Authorized) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	return nil
}
