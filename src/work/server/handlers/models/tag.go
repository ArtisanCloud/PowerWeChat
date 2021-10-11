package models

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
)

const CALLBACK_EVENT_CHANGE_TYPE_UPDATE_TAG   = "update_tag"

type EventTagUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	TagId         string   `xml:"TagId"`
	AddUserItems  string   `xml:"AddUserItems"`
	DelUserItems  string   `xml:"DelUserItems"`
	AddPartyItems string   `xml:"AddPartyItems"`
	DelPartyItems string   `xml:"DelPartyItems"`
}

