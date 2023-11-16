package server

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/response"
	openplatform "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/server/callbacks"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/server/handlers"
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
	guard.OverrideNotify()

	return guard

}

func (guard *Guard) Notify(request *http.Request, closure func(content *openplatform.Callback, decrypted []byte, infoType string) interface{}) (httpRS *http.Response, err error) {
	// validate the signature
	_, err = guard.Validate(request)
	if err != nil {
		return nil, err
	}

	// read body content
	requestXML, _ := io.ReadAll(request.Body)
	request.Body = io.NopCloser(bytes.NewBuffer(requestXML))
	println(string(requestXML))

	// convert to openplatform event
	callbackEvent := &openplatform.Callback{}
	err = xml.Unmarshal(requestXML, callbackEvent)
	if err != nil {
		return nil, err
	}

	// decrypt event content
	var bufDecrypted []byte = nil
	if guard.IsSafeMode(request) && callbackEvent.Encrypt != "" {
		bufDecrypted, err = guard.DecryptEvent(callbackEvent.Encrypt)
	}

	// call the closure for handling the event
	msg := &response2.ResponseVerifyTicket{}
	err = xml.Unmarshal(bufDecrypted, msg)
	result := closure(callbackEvent, bufDecrypted, msg.InfoType)

	// convert the result to http response
	var buffResult []byte
	// 如果是字符串， "success", "failed"
	if str, ok := result.(string); ok {
		buffResult = []byte(str)
	} else {
		strResult, err := object.JsonEncode(result)
		if err != nil {
			return nil, err
		}
		buffResult = []byte(strResult)
	}

	httpRS = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(buffResult)),
	}

	return httpRS, err

}

func (guard *Guard) DecryptEvent(content string) (bufDecrypted []byte, err error) {

	encryptor := (*guard.App).GetComponent("Encryptor").(*kernel.Encryptor)

	bufDecrypted, cryptErr := encryptor.DecryptContent(content)
	if cryptErr != nil {
		return nil, errors.New(cryptErr.ErrMsg)
	}

	return bufDecrypted, err

}

// Override Validate
func (guard *Guard) OverrideResolve() {
	guard.Resolve = func(request *http.Request) (httpRS *http.Response, err error) {
		guard.registerHandlers()

		message, err := guard.GetMessage(request)

		if message.InfoType != "" {
			_ = guard.Dispatch(request, GetOpenPlatformEvent(message.InfoType), nil, message)
		}
		httpRS = &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(kernel.SUCCESS_EMPTY_RESPONSE))),
		}

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
		request.Body = io.NopCloser(bytes.NewBuffer(b))
	}
	verifyTicket = &response2.ResponseVerifyTicket{}
	err = guard.parseMessage(string(b), verifyTicket)
	if err != nil {
		return nil, err
	}

	return verifyTicket, err

}

func (guard *Guard) parseMessage(content string, callback interface{}) (err error) {

	if len(content) == 0 {
		return errors.New("empty content")
	}
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

	return nil
}
