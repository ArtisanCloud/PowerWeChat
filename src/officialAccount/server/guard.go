package server

import (
	"bytes"
	logger2 "github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"io/ioutil"
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

	guard.OverrideShouldReturnRawResponse()

	return guard

}

func (guard *Guard) VerifyURL(request *http.Request) (httpRS *http.Response, err error) {
	logger := (*guard.App).GetComponent("Logger").(*logger2.Logger)

	_, err = guard.Validate(request)
	if err != nil {
		return nil, err
	}

	bodyData := ioutil.NopCloser(bytes.NewBufferString(request.URL.Query().Get("echostr")))
	rs := &http.Response{
		Body:       bodyData,
		StatusCode: http.StatusOK,
	}

	logger.Info("Server response created:", "content", request.URL.Query().Get("echostr"))

	return rs, err
}

// Override Validate
func (guard *Guard) OverrideShouldReturnRawResponse() {
	guard.ShouldReturnRawResponse = func(request *http.Request) bool {
		if request == nil || request.URL.Query().Get("echostr") == "" {
			return false
		}
		return true
	}
}
