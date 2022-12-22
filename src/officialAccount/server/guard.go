package server

import (
	"bytes"
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	logger2 "github.com/ArtisanCloud/PowerLibs/v2/logger"
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

func (guard *Guard) VerifyURL(request *http.Request) (httpRS *response.HttpResponse, err error) {
	logger := (*guard.App).GetComponent("Logger").(*logger2.Logger)

	_, err = guard.Validate(request)
	if err != nil {
		return nil, err
	}

	rs := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(request.URL.Query().Get("echostr"))),
		StatusCode: http.StatusOK,
	}

	httpRS = response.NewHttpResponse(http.StatusOK)
	httpRS.Response = rs

	logger.Info("Server response created:", "content", httpRS.GetBody())

	return httpRS, err
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
