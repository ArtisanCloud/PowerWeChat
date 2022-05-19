package server

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
)

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		kernel.NewServerGuard(app),
	}

	guard.OverrideShouldReturnRawResponse()

	return guard

}

// Override Validate
func (guard *Guard) OverrideShouldReturnRawResponse() {
	guard.ShouldReturnRawResponse = func() bool {
		request := (*guard.App).GetExternalRequest()
		if request == nil || request.URL.Query().Get("echostr") == "" {
			return false
		}
		return true
	}
}
