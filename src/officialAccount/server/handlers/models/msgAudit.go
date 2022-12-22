package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

const CALLBACK_EVENT_MSGAUDIT_NOTIFY = "msgaudit_notify"

type EventMsgAuditNotify struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}
