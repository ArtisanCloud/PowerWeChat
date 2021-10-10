package server

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		kernel.NewServerGuard(app),
	}

	guard.OverrideIsSafeMode()
	guard.OverrideValidate()
	guard.OverrideShouldReturnRawResponse()

	return guard

}

// Override Validate
func (guard *Guard) OverrideIsSafeMode() {
	guard.IsSafeMode = func() bool {
		return true
	}
}

func (guard *Guard) OverrideValidate() {
	guard.Validate = func() (*kernel.ServerGuard, error) {
		return guard.ServerGuard, nil
	}
}

func (guard *Guard) OverrideShouldReturnRawResponse() {
	guard.ShouldReturnRawResponse = func() bool {
		request := guard.ExternalRequest
		if request == nil || request.URL.Query().Get("echostr") == "" {
			return false
		}
		return true
	}
}
