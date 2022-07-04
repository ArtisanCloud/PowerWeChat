package server

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/server/handlers"
	"io"
	"io/ioutil"
	"net/http"
)

const EVENT_AUTHORIZED int = 12000              // "authorized"
const EVENT_UNAUTHORIZED int = 12002            // "unauthorized"
const EVENT_UPDATE_AUTHORIZED int = 12003       //  "updateauthorized"
const EVENT_COMPONENT_VERIFY_TICKET int = 12004 // "component_verify_ticket"
const EVENT_THIRD_FAST_REGISTERED int = 120005  //  "notify_third_fasteregister"

func GetOpenPlatformEvent(infoType string) int {
	switch infoType {
	case "authorized":
		return EVENT_AUTHORIZED
	case "unauthorized":
		return EVENT_UNAUTHORIZED
	case "updateauthorized":
		return EVENT_UPDATE_AUTHORIZED
	case "component_verify_ticket":
		return EVENT_COMPONENT_VERIFY_TICKET
	case "notify_third_fasteregister":
		return EVENT_THIRD_FAST_REGISTERED
	}

	return -1
}

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		kernel.NewServerGuard(app),
	}

	guard.OverrideResolve()

	return guard

}

// Override Validate
func (guard *Guard) OverrideResolve() {
	guard.Resolve = func(request *http.Request) (httpRS *response.HttpResponse, err error) {
		guard.registerHandlers()

		message, err := guard.GetMessage(request)

		var rs *http.Response = &http.Response{}
		if message.InfoType != "" {
			_ = guard.Dispatch(request, GetOpenPlatformEvent(message.InfoType), nil, message)
		}
		httpRS = response.NewHttpResponse(http.StatusOK)
		rs.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(kernel.SUCCESS_EMPTY_RESPONSE)))
		httpRS.Response = rs

		return httpRS, nil
	}
}

func (guard *Guard) registerHandlers() {

	guard.On(EVENT_AUTHORIZED, handlers.NewAuthorized(guard.App))
	guard.On(EVENT_UNAUTHORIZED, handlers.NewUnauthorized(guard.App))
	guard.On(EVENT_UPDATE_AUTHORIZED, handlers.NewUpdateAuthorized(guard.App))
	guard.On(EVENT_COMPONENT_VERIFY_TICKET, handlers.NewVerifyTicketRefreshed(guard.App))

}

func (guard *Guard) GetMessage(request *http.Request) (verifyTicket *response2.ResponseVerifyTicket, err error) {
	var b = []byte("")
	if request.Body != http.NoBody {
		b, err = io.ReadAll(request.Body)
		if err != nil || b == nil {
			return nil, err
		}
	}
	verifyTicket = &response2.ResponseVerifyTicket{}
	err = guard.parseMessage(string(b), verifyTicket)
	if err != nil {
		return nil, err
	}

	return verifyTicket, err

}

func (guard *Guard) parseMessage(content string, callback interface{}) (err error) {

	if len(content) > 0 {
		if content[0:1] == "<" {
			err = xml.Unmarshal([]byte(content), callback)
			if err != nil {
				return err
			}
		} else {
			// Handle JSON format.
			err = object.JsonDecode([]byte(content), callback)
			if err != nil {
				return err
			}
		}
	} else {
		errors.New("empty content")
	}

	return err
}
