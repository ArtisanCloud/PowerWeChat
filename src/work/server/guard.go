package server

import (
	"encoding/xml"
	"errors"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	models2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/server/handlers/models"
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

	guard.OverrideIsSafeMode()
	guard.OverrideValidate()
	guard.OverrideShouldReturnRawResponse()
	guard.OverrideToCallbackType()

	return guard

}

// Override Validate
func (guard *Guard) OverrideIsSafeMode() {
	guard.IsSafeMode = func(request *http.Request) bool {
		return true
	}
}

func (guard *Guard) OverrideValidate() {
	guard.Validate = func(request *http.Request) (*kernel.ServerGuard, error) {
		return guard.ServerGuard, nil
	}
}

func (guard *Guard) OverrideShouldReturnRawResponse() {
	guard.ShouldReturnRawResponse = func(request *http.Request) bool {
		if request == nil || request.URL.Query().Get("echostr") == "" {
			return false
		}
		return true
	}
}

func (guard *Guard) OverrideToCallbackType() {

	// switch message type
	guard.ToCallbackType = func(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error) {
		switch callbackHeader.GetMsgType() {

		// msg type is message
		case models.CALLBACK_MSG_TYPE_TEXT:
			decryptMsg := models2.MessageText{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case models.CALLBACK_MSG_TYPE_IMAGE:
			decryptMsg := models2.MessageImage{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case models.CALLBACK_MSG_TYPE_VOICE:
			decryptMsg := models2.MessageVoice{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case models.CALLBACK_MSG_TYPE_VIDEO:
			decryptMsg := models2.MessageVideo{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case models.CALLBACK_MSG_TYPE_LOCATION:
			decryptMsg := models2.MessageLocation{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		case models.CALLBACK_MSG_TYPE_LINK:
			decryptMsg := models2.MessageLink{}
			err = xml.Unmarshal(buf, &decryptMsg)
			decryptMessage = decryptMsg
			break

		// msg type is event
		case models.CALLBACK_MSG_TYPE_EVENT:
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
	case models2.CALLBACK_EVENT_CHANGE_CONTACT:
		decryptMessage, err = guard.toCallbackEventChangeType(callbackHeader, buf)
		break

	// events
	case models2.CALLBACK_EVENT_SUBSCRIBE:
		decryptMsg := &models2.EventSubscribe{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_ENTER_AGENT:
		decryptMsg := &models2.EventEnterAgent{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_LOCATION:
		decryptMsg := &models2.EventLocation{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_BATCH_JOB_RESULT:
		decryptMsg := &models2.EventBatchJobResult{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_CLICK:
		decryptMsg := &models2.EventClick{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_VIEW:
		decryptMsg := &models2.EventView{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_SCANCODE_PUSH:
		decryptMsg := &models2.EventScanCodePush{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_SCANCODE_WAITMSG:
		decryptMsg := &models2.EventScancodeWaitMsg{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_PIC_SYSPHOTO:
		decryptMsg := &models2.EventPicSysPhoto{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM:
		decryptMsg := &models2.EventPicPhotoOrAlbum{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_PIC_WEIXIN:
		decryptMsg := &models2.EventPicWeixin{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_LOCATION_SELECT:
		decryptMsg := &models2.EventLocationSelect{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_OPEN_APPROVAL_CHANGE:
		decryptMsg := &models2.EventOpenApprovalChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_SHARE_AGENT_CHANGE:
		decryptMsg := &models2.EventShareAgentChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT:
		decryptMsg := &models2.EventTemplateCardMenuEvent{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_LIVING_STATUS_CHANGE:
		decryptMsg := &models2.EventLivingStatusChange{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_MSGAUDIT_NOTIFY:
		decryptMsg := &models2.EventMsgAuditNotify{}
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

	// change type is for user event
	case models2.CALLBACK_EVENT_CHANGE_TYPE_CREATE_USER:
		decryptMsg := &models2.EventUserCreate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_USER:
		decryptMsg := &models2.EventUserUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_DELETE_USER:
		decryptMsg := &models2.EventUserDelete{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	// change type is for party event
	case models2.CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY:
		decryptMsg := &models2.EventPartyCreate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_PARTY:
		decryptMsg := &models2.EventPartyUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_DELETE_PARTY:
		decryptMsg := &models2.EventPartyDelete{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	// change type is for tag event
	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_TAG:
		decryptMsg := &models2.EventTagUpdate{}
		err = xml.Unmarshal(buf, decryptMsg)
		decryptMessage = decryptMsg
		break

	default:
		return nil, errors.New("not found wechat event")
	}

	return decryptMessage, err
}
