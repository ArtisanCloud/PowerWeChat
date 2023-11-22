package server

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	kernelModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/handlers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/models"
)

const (
	EVENT_SUITE_TICKET         int = 52000 // "suite_ticket"
	EVENT_CREATE_AUTH          int = 52100 // "create_auth"
	EVENT_CHANGE_AUTH          int = 52101 //  "change_auth"
	EVENT_CANCEL_AUTH          int = 52102 //  "cancel_auth"
	EVENT_CREATE_USER          int = 52200 // "change_contact->create_user"
	EVENT_UPDATE_USER          int = 52201 // "change_contact->update_user"
	EVENT_DELETE_USER          int = 52202 // "change_contact->delete_user"
	EVENT_CREATE_PARTY         int = 52210 // "change_contact->create_party"
	EVENT_UPDATE_PARTY         int = 52211 // "change_contact->update_party"
	EVENT_DELETE_PARTY         int = 52212 // "change_contact->delete_party"
	EVENT_UPDATE_TAG           int = 52220 // "change_contact->update_tag"
	EVENT_SHARE_AGENT_CHANGE   int = 52300 // "share_agent_change"
	EVENT_SHARE_CHAIN_CHANGE   int = 52310 // "share_chain_change"
	EVENT_RESET_PERMANENT_CODE int = 52400 // "reset_permanent_code"
	EVENT_CORP_ARCH_AUTH       int = 52600 // "corp_arch_auth"
	EVENT_APPROVE_SPECIAL_AUTH int = 52700 // "approve_special_auth"
	EVENT_CANCEL_SPECIAL_AUTH  int = 52701 // "cancel_special_auth"
)

func GetEventType(infoType models.InfoType, changeType models.ChangeType) int {
	switch infoType {
	case models.InfoTypeSuiteTicket:
		return EVENT_SUITE_TICKET
	case models.InfoTypeCreateAuth:
		return EVENT_CREATE_AUTH
	case models.InfoTypeChangeAuth:
		return EVENT_CHANGE_AUTH
	case models.InfoTypeCancelAuth:
		return EVENT_CANCEL_AUTH
	case models.InfoTypeChangeContact:
		switch changeType {
		case models.ChangeTypeCreateUser:
			return EVENT_CREATE_USER
		case models.ChangeTypeUpdateUser:
			return EVENT_UPDATE_USER
		case models.ChangeTypeDeleteUser:
			return EVENT_DELETE_USER
		case models.ChangeTypeCreateParty:
			return EVENT_CREATE_PARTY
		case models.ChangeTypeUpdateParty:
			return EVENT_UPDATE_USER
		case models.ChangeTypeDeleteParty:
			return EVENT_DELETE_USER
		case models.ChangeTypeUpdateTag:
			return EVENT_UPDATE_TAG
		}
		return -1
	case models.InfoTypeShareAgentChange:
		return EVENT_SHARE_AGENT_CHANGE
	case models.InfoTypeShareChainChange:
		return EVENT_SHARE_CHAIN_CHANGE
	case models.InfoTypeResetPermanentCode:
		return EVENT_RESET_PERMANENT_CODE
	case models.InfoTypeCorpArchAuth:
		return EVENT_CORP_ARCH_AUTH
	case models.InfoTypeApproveSpecialAuth:
		return EVENT_APPROVE_SPECIAL_AUTH
	case models.InfoTypeCancelSpecialAuth:
		return EVENT_CANCEL_SPECIAL_AUTH
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
	guard.OverrideIsSafeMode()

	return guard

}

// Override Validate
func (guard *Guard) OverrideIsSafeMode() {
	guard.IsSafeMode = func(request *http.Request) bool {
		return true
	}
}

func (guard *Guard) Notify(request *http.Request, closure func(content *kernelModels.Callback, ev models.IEvent) interface{}) (httpRS *http.Response, err error) {
	// validate the signature
	_, err = guard.Validate(request)
	if err != nil {
		return nil, err
	}

	// read body content
	requestXML, _ := io.ReadAll(request.Body)
	request.Body = io.NopCloser(bytes.NewBuffer(requestXML))
	println(string(requestXML))

	var callbackEvent kernelModels.Callback
	err = xml.Unmarshal(requestXML, &callbackEvent)
	if err != nil {
		return nil, err
	}

	// decrypt event content
	var bufDecrypted []byte = nil
	if guard.IsSafeMode(request) && callbackEvent.Encrypt != "" {
		bufDecrypted, err = guard.DecryptEvent(callbackEvent.Encrypt)
	}

	// call the closure for handling the event
	ev, err := models.DecodeEvent(bufDecrypted)
	if err != nil {
		return nil, err
	}

	result := closure(&callbackEvent, ev)

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

		if message, err := guard.GetMessage(request); err != nil {
			return nil, err
		} else if infoType := message.GetInfoType(); infoType != "" {
			_ = guard.Dispatch(request, GetEventType(infoType, message.GetChangeType()), nil, message)
		}
		httpRS = &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(kernel.SUCCESS_EMPTY_RESPONSE))),
		}

		return httpRS, nil
	}
}

func (guard *Guard) registerHandlers() {

	// guard.On(EVENT_AUTHORIZED, handlers.NewAuthorized(guard.App))
	// guard.On(EVENT_UNAUTHORIZED, handlers.NewUnauthorized(guard.App))
	// guard.On(EVENT_UPDATE_AUTHORIZED, handlers.NewUpdateAuthorized(guard.App))
	guard.On(EVENT_SUITE_TICKET, handlers.NewSuiteTicket(guard.App))

}

func (guard *Guard) GetMessage(request *http.Request) (ev models.IEvent, err error) {
	var b = []byte("")
	if request.Body != http.NoBody {
		b, err = io.ReadAll(request.Body)
		if err != nil || b == nil {
			return nil, err
		}
		request.Body = io.NopCloser(bytes.NewBuffer(b))
	}
	baseEv := new(models.BaseEvent)
	if err = guard.parseMessage(string(b), baseEv); err != nil {
		return nil, err
	}
	if ev, err = baseEv.ToEvent(); err != nil {
		return nil, err
	}
	if err = guard.parseMessage(string(b), ev); err != nil {
		return nil, err
	}
	return ev, err

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
