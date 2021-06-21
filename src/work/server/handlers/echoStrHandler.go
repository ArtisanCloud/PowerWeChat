package handlers

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"net/http"
)

type EchoStrHandler struct {
	contract.EventHandlerInterface

	App     *kernel.ApplicationInterface
	Request *http.Request
}

func NewBaseClient(app *ApplicationInterface) *EchoStrHandler {
	handler := &EchoStrHandler{
		App: app,
	}

	return handler
}

func (handler *EchoStrHandler) handle(payload interface{}) {
	descrypted := (*handler.App).
}
