package server

import (
	"encoding/xml"
	"errors"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
	models2 "github.com/ArtisanCloud/power-wechat/src/work/server/handlers/models"
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
		request := (*guard.App).GetExternalRequest()
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
			decryptMessage = models2.MessageText{}
			break

		case models.CALLBACK_MSG_TYPE_IMAGE:
			decryptMessage = models2.MessageImage{}
			break

		case models.CALLBACK_MSG_TYPE_VOICE:
			decryptMessage = models2.MessageVoice{}
			break

		case models.CALLBACK_MSG_TYPE_VIDEO:
			decryptMessage = models2.MessageVideo{}
			break

		case models.CALLBACK_MSG_TYPE_LOCATION:
			decryptMessage = models2.MessageLocation{}
			break

		case models.CALLBACK_MSG_TYPE_LINK:
			decryptMessage = models2.MessageLink{}
			break

		// msg type is event
		case models.CALLBACK_MSG_TYPE_EVENT:
			decryptMessage, err = guard.toCallbackEvent(callbackHeader, buf)
			return decryptMessage, err

		default:
			return nil, errors.New("not found wechat msg type")
		}

		err = xml.Unmarshal(buf, &decryptMessage)

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
		decryptMessage = &models2.EventSubscribe{}
		break

	case models2.CALLBACK_EVENT_ENTER_AGENT:
		decryptMessage = &models2.EventEnterAgent{}
		break

	case models2.CALLBACK_EVENT_LOCATION:
		decryptMessage = &models2.EventLocation{}
		break

	case models2.CALLBACK_EVENT_BATCH_JOB_RESULT:
		decryptMessage = &models2.EventBatchJobResult{}
		break

	case models2.CALLBACK_EVENT_CLICK:
		decryptMessage = &models2.EventClick{}
		break

	case models2.CALLBACK_EVENT_VIEW:
		decryptMessage = &models2.EventView{}
		break

	case models2.CALLBACK_EVENT_SCANCODE_PUSH:
		decryptMessage = &models2.EventScanCodePush{}
		break

	case models2.CALLBACK_EVENT_SCANCODE_WAITMSG:
		decryptMessage = &models2.EventScancodeWaitMsg{}
		break

	case models2.CALLBACK_EVENT_PIC_SYSPHOTO:
		decryptMessage = &models2.EventPicSysPhoto{}
		break

	case models2.CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM:
		decryptMessage = &models2.EventPicPhotoOrAlbum{}
		break

	case models2.CALLBACK_EVENT_PIC_WEIXIN:
		decryptMessage = &models2.EventPicWeixin{}
		break

	case models2.CALLBACK_EVENT_LOCATION_SELECT:
		decryptMessage = &models2.EventLocationSelect{}
		break

	case models2.CALLBACK_EVENT_OPEN_APPROVAL_CHANGE:
		decryptMessage = &models2.EventOpenApprovalChange{}
		break

	case models2.CALLBACK_EVENT_SHARE_AGENT_CHANGE:
		decryptMessage = &models2.EventShareAgentChange{}
		break

	case models2.CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT:
		decryptMessage = &models2.EventTemplateCardMenuEvent{}
		break

	case models2.CALLBACK_EVENT_LIVING_STATUS_CHANGE:
		decryptMessage = &models2.EventLivingStatusChange{}
		break

	case models2.CALLBACK_EVENT_MSGAUDIT_NOTIFY:
		decryptMessage = &models2.EventMsgAuditNotify{}
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
		decryptMessage = &models2.EventUserCreate{}
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_USER:
		decryptMessage = &models2.EventUserUpdate{}
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_DELETE_USER:
		decryptMessage = &models2.EventUserDelete{}
		break

	// change type is for party event
	case models2.CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY:
		decryptMessage = &models2.EventPartyCreate{}
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_PARTY:
		decryptMessage = &models2.EventPartyUpdate{}
		break

	case models2.CALLBACK_EVENT_CHANGE_TYPE_DELETE_PARTY:
		decryptMessage = &models2.EventPartyDelete{}
		break

	// change type is for tag event
	case models2.CALLBACK_EVENT_CHANGE_TYPE_UPDATE_TAG:
		decryptMessage = &models2.EventTagUpdate{}
		break

	default:
		return nil, errors.New("not found wechat event")
	}

	err = xml.Unmarshal(buf, decryptMessage)

	return decryptMessage, err
}
