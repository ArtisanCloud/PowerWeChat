package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

const CALLBACK_EVENT_LIVING_STATUS_CHANGE = "living_status_change"

type EventLivingStatusChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	LivingID string `xml:"LivingId"`
	Status   string `xml:"Status"`
	AgentID  string `xml:"AgentID"`
}
