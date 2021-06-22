package server

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"net/http"
)

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		kernel.NewServerGuard(app),
	}
	return guard

}

func (guard *Guard) validate() *Guard {
	return guard
}

func (guard *Guard) isSafeMode() bool {
	return true
}

func (guard *Guard) shouldReturnRawResponse() bool {
	request := (*guard.App).GetComponent("ExternalRequest").(*http.Request)
	if request == nil || request.URL.Query().Get("echostr") == "" {
		return false
	}
	return true
}
