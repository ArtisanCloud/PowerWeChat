package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

type MessageText struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Content string `xml:"Content"`
	MsgID   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageImage struct {
	contract.EventInterface
	models.CallbackMessageHeader
	PicUrl  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`
	MsgID   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageVoice struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaID string `xml:"MediaId"`
	Format  string `xml:"Format"`
	MsgID   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageLocation struct {
	contract.EventInterface
	models.CallbackMessageHeader
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
	MsgID     string `xml:"MsgId"`
	AgentID   string `xml:"AgentID"`
	AppType   string `xml:"AppType"`
}

type MessageVideo struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
	MsgID        string `xml:"MsgId"`
	AgentID      string `xml:"AgentID"`
}

type MessageLink struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
	PicUrl      string `xml:"PicUrl"`
	MsgID       string `xml:"MsgId"`
	AgentID     string `xml:"AgentID"`
}
