package server

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	kernelModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/handlers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/models"
	workModels "github.com/ArtisanCloud/PowerWeChat/v3/src/work/server/handlers/models"
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
	// guard.OverrideNotify()
	guard.OverrideIsSafeMode()
	guard.OverrideToCallbackType()

	return guard

}

// Override Validate
func (guard *Guard) OverrideIsSafeMode() {
	guard.IsSafeMode = func(request *http.Request) bool {
		return true
	}
}

func (guard *Guard) Notify(request *http.Request, closure func(content *kernelModels.Callback, ev models.IEvent, callbackMsg interface{}) interface{}) (httpRS *http.Response, err error) {
	// validate the signature
	_, err = guard.Validate(request)
	if err != nil {
		return nil, err
	}

	var (
		buffResult  []byte
		contentType string
		echoStr     = request.URL.Query().Get("echostr")
	)
	if request.Body == http.NoBody && echoStr != "" {
		buffResult, err = guard.DecryptEvent(echoStr)
		if err != nil {
			return nil, err
		}
		contentType = "html/text"
	} else if request.Body != http.NoBody {
		var (
			callbackEvent kernelModels.Callback
			ev            models.IEvent
		)
		// read body content
		var buf bytes.Buffer
		tee := io.TeeReader(request.Body, &buf)
		// requestXML, _ := io.ReadAll(request.Body)
		request.Body = io.NopCloser(&buf)
		// println(string(requestXML))

		// err = xml.Unmarshal(requestXML, &callbackEvent)
		err = xml.NewDecoder(tee).Decode(&callbackEvent)
		if err != nil {
			return nil, err
		}
		// decrypt event content
		var bufDecrypted []byte = nil
		if guard.IsSafeMode(request) && callbackEvent.Encrypt != "" {
			bufDecrypted, err = guard.DecryptEvent(callbackEvent.Encrypt)
		}
		var callbackMsg interface{}
		// call the closure for handling the event
		ev, err = models.DecodeEvent(bufDecrypted)
		if err != nil {
			if callbackMsg, err = guard.ToCallbackType(ev, bufDecrypted); err != nil {
				return nil, err
			}
		}
		result := closure(&callbackEvent, ev, callbackMsg)

		// convert the result to http response
		// 如果是字符串， "success", "failed"
		if str, ok := result.(string); ok {
			buffResult = []byte(str)
			contentType = "html/text"
		} else {
			strResult, err := object.JsonEncode(result)
			if err != nil {
				return nil, err
			}
			contentType = "application/json"
			buffResult = []byte(strResult)
		}
	} else {
		return nil, errors.New("missing request body")
	}

	httpRS = &http.Response{
		StatusCode:    http.StatusOK,
		ContentLength: int64(len(buffResult)),
		Header: http.Header{
			"Content-Type": []string{contentType},
		},
		Body: io.NopCloser(bytes.NewBuffer(buffResult)),
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

		if ev, _, err := guard.GetMessage(request); err != nil {
			return nil, err
		} else if infoType := ev.GetInfoType(); infoType != "" {
			_ = guard.Dispatch(request, GetEventType(infoType, ev.GetChangeType()), nil, ev)
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

func (guard *Guard) GetMessage(request *http.Request) (ev models.IEvent, msg interface{}, err error) {
	var b = []byte("")
	if request.Body != http.NoBody {
		b, err = io.ReadAll(request.Body)
		if err != nil || b == nil {
			return nil, nil, err
		}
		request.Body = io.NopCloser(bytes.NewBuffer(b))
	}
	baseEv := new(models.BaseEvent)
	if err = guard.parseMessage(string(b), baseEv); err != nil {
		return nil, nil, err
	}
	if ev, err = baseEv.ToEvent(); err != nil {
		if msg, err = guard.ToCallbackType(baseEv.CallbackMessageHeader, b); err != nil {
			return nil, nil, err
		} else {
			return nil, msg, nil
		}
	}
	if err = guard.parseMessage(string(b), ev); err != nil {
		return nil, nil, err
	}
	return ev, nil, err
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

func (guard *Guard) OverrideToCallbackType() {

	guard.ToCallbackType = func(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error) {
		switch callbackHeader.GetMsgType() {

		// msg type is message
		case kernelModels.CALLBACK_MSG_TYPE_TEXT:
			decryptMsg := workModels.MessageText{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case kernelModels.CALLBACK_MSG_TYPE_IMAGE:
			decryptMsg := workModels.MessageImage{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case kernelModels.CALLBACK_MSG_TYPE_VOICE:
			decryptMsg := workModels.MessageVoice{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case kernelModels.CALLBACK_MSG_TYPE_VIDEO:
			decryptMsg := workModels.MessageVideo{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case kernelModels.CALLBACK_MSG_TYPE_LOCATION:
			decryptMsg := workModels.MessageLocation{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case kernelModels.CALLBACK_MSG_TYPE_LINK:
			decryptMsg := workModels.MessageLink{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		// msg type is event
		case kernelModels.CALLBACK_MSG_TYPE_EVENT:
			decryptMessage, err = guard.toCallbackEvent(callbackHeader, buf)
			return decryptMessage, err

		default:
			return nil, errors.New("not found wechat msg type")
		}

		return decryptMessage, err
	}

}

// switch event
func (guard *Guard) toCallbackEvent(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error) {

	switch callbackHeader.GetEvent() {

	// event is change contact
	case workModels.CALLBACK_EVENT_CUSTOMER_ACQUISITION,
		workModels.CALLBACK_EVENT_CHANGE_EXTERNAL_CONTACT,
		workModels.CALLBACK_EVENT_CHANGE_CONTACT:
		decryptMessage, err = guard.toCallbackEventChangeType(callbackHeader, buf)
		break
	case workModels.CALLBACK_EVENT_CHANGE_EXTERNAL_CHAT:
		decryptMessage, err = guard.toCallbackChatEventChangeType(callbackHeader, buf)
		break

	// events
	case workModels.CALLBACK_EVENT_SUBSCRIBE:
		decryptMsg := &workModels.EventSubscribe{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_ENTER_AGENT:
		decryptMsg := &workModels.EventEnterAgent{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_LOCATION:
		decryptMsg := &workModels.EventLocation{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_BATCH_JOB_RESULT:
		decryptMsg := &workModels.EventBatchJobResult{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CLICK:
		decryptMsg := &workModels.EventClick{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_VIEW:
		decryptMsg := &workModels.EventView{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_SCANCODE_PUSH:
		decryptMsg := &workModels.EventScanCodePush{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_SCANCODE_WAITMSG:
		decryptMsg := &workModels.EventScancodeWaitMsg{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_PIC_SYSPHOTO:
		decryptMsg := &workModels.EventPicSysPhoto{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM:
		decryptMsg := &workModels.EventPicPhotoOrAlbum{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_PIC_WEIXIN:
		decryptMsg := &workModels.EventPicWeixin{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_LOCATION_SELECT:
		decryptMsg := &workModels.EventLocationSelect{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_OPEN_APPROVAL_CHANGE:
		decryptMsg := &workModels.EventOpenApprovalChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_SHARE_AGENT_CHANGE:
		decryptMsg := &workModels.EventShareAgentChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT:
		decryptMsg := &workModels.EventTemplateCardMenuEvent{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_LIVING_STATUS_CHANGE:
		decryptMsg := &workModels.EventLivingStatusChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_MSGAUDIT_NOTIFY:
		decryptMsg := &workModels.EventMsgAuditNotify{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	default:
		return nil, errors.New("not found wechat event")
	}

	return decryptMessage, err
}

// switch event change type
func (guard *Guard) toCallbackEventChangeType(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error) {

	switch callbackHeader.GetChangeType() {

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_OPEN_PROFILE:
		decryptMsg := &workModels.EventCustomerAcquisitionOpenProfile{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_FRIEND_REQUEST:
		decryptMsg := &workModels.EventCustomerAcquisitionFriendRequest{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_ADD_EXTERNAL_CONTACT:
		decryptMsg := &workModels.EventExternalUserAdd{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_ADD_HALF_EXTERNAL_CONTACT:
		decryptMsg := &workModels.EventExternalUserAddHalf{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_DEL_EXTERNAL_CONTACT:
		decryptMsg := &workModels.EventExternalUserEdit{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	// change type is for user event
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_CREATE_USER:
		decryptMsg := &workModels.EventUserCreate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_USER:
		decryptMsg := &workModels.EventUserUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_DELETE_USER:
		decryptMsg := &workModels.EventUserDelete{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	// change type is for party event
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY:
		decryptMsg := &workModels.EventPartyCreate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_PARTY:
		decryptMsg := &workModels.EventPartyUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case workModels.CALLBACK_EVENT_CHANGE_TYPE_DELETE_PARTY:
		decryptMsg := &workModels.EventPartyDelete{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	// change type is for tag event
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_TAG:
		decryptMsg := &workModels.EventTagUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	default:
		return nil, errors.New("not found wechat event")
	}

	return decryptMessage, err
}

// switch chat event change type
func (guard *Guard) toCallbackChatEventChangeType(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error) {

	switch callbackHeader.GetChangeType() {
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_CREATE:
		decryptMsg := &workModels.EventExternalChatCreate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_UPDATE:
		decryptMsg := &workModels.EventExternalChatUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break
	case workModels.CALLBACK_EVENT_CHANGE_TYPE_DISMISS:
		decryptMsg := &workModels.EventExternalChatDismiss{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break
	default:
		return nil, errors.New("not found wechat event")
	}

	return decryptMessage, err
}
