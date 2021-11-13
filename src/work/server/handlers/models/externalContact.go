package models

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/models"
)

const (
	CALLBACK_EVENT_CHANGE_EXTERNAL_CONTACT   = "change_external_contact"
	CALLBACK_EVENT_CHANGE_EXTERNAL_TAG   = "change_external_tag"

	CALLBACK_EVENT_CHANGE_TYPE_ADD_EXTERNAL_CONTACT  = "add_external_contact"
	CALLBACK_EVENT_CHANGE_TYPE_EDIT_EXTERNAL_CONTACT  = "edit_external_contact"
	CALLBACK_EVENT_CHANGE_TYPE_ADD_HALF_EXTERNAL_CONTACT  = "add_half_external_contact"
	CALLBACK_EVENT_CHANGE_TYPE_DEL_EXTERNAL_CONTACT  = "del_external_contact"
	CALLBACK_EVENT_CHANGE_TYPE_DEL_FOLLOW_USER  = "del_follow_user"
	CALLBACK_EVENT_CHANGE_TYPE_TRANSFER_FAIL  = "transfer_fail"
	CALLBACK_EVENT_CHANGE_TYPE_CREATE  = "create"
	CALLBACK_EVENT_CHANGE_TYPE_UPDATE  = "update"
	CALLBACK_EVENT_CHANGE_TYPE_DISMISS  = "dismiss"
	CALLBACK_EVENT_CHANGE_TYPE_DELETE  = "delete"
	CALLBACK_EVENT_CHANGE_TYPE_SHUFFLE  = "shuffle"

	CALLBACK_EVENT_UPDATE_DETAIL = "add_member"
	CALLBACK_EVENT_TAG_TYPE = "tag"

)


type EventExternalUserAdd struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
	State          string   `xml:"State"`
	WelcomeCode    string   `xml:"WelcomeCode"`
}


type EventExternalUserEdit struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
}

type EventExternalUserAddHalf struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
	State          string   `xml:"State"`
	WelcomeCode    string   `xml:"WelcomeCode"`
}


type EventExternalUserDel struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
	Source         string   `xml:"Source"`
}

type EventExternalUserDelFollowUser struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
}

type EventExternalUserUpdateAddMember struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatId       string   `xml:"ChatId"`
	ChangeType   string   `xml:"ChangeType"`
	UpdateDetail string   `xml:"UpdateDetail"`
	JoinScene    string   `xml:"JoinScene"`
	QuitScene    string   `xml:"QuitScene"`
	MemChangeCnt string   `xml:"MemChangeCnt"`
}


type EventExternalUserDismiss struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatId       string   `xml:"ChatId"`
}


type EventExternalUserTagCreate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID           string   `xml:"Id"`
	TagType      string   `xml:"TagType"`
	StrategyId   string   `xml:"StrategyId"`
}


type EventExternalUserTagUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID           string   `xml:"Id"`
	TagType      string   `xml:"TagType"`
	StrategyId   string   `xml:"StrategyId"`
}


type EventExternalUserTagDelete struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID           string   `xml:"Id"`
	TagType      string   `xml:"TagType"`
	StrategyId   string   `xml:"StrategyId"`
}


type EventExternalUserTagShuffle struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID           string   `xml:"Id"`
	StrategyId   string   `xml:"StrategyId"`
	ChangeType   string   `xml:"ChangeType"`
}
