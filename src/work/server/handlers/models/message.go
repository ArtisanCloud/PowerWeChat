package models

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
)

type MessageText struct {
	models.CallbackMessageHeader
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	AgentID      string   `xml:"AgentID"`
}

