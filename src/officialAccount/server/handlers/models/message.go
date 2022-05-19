package models

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/models"
)

type MessageText struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Content string `xml:"Content"`
	MsgId   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageImage struct {
	contract.EventInterface
	models.CallbackMessageHeader
	PicUrl  string `xml:"PicUrl"`
	MediaId string `xml:"MediaId"`
	MsgId   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageVoice struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaId string `xml:"MediaId"`
	Format  string `xml:"Format"`
	MsgId   string `xml:"MsgId"`
	AgentID string `xml:"AgentID"`
}

type MessageLocation struct {
	contract.EventInterface
	models.CallbackMessageHeader
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
	MsgId     string `xml:"MsgId"`
	AgentID   string `xml:"AgentID"`
	AppType   string `xml:"AppType"`
}

type MessageVideo struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        string `xml:"MsgId"`
	AgentID      string `xml:"AgentID"`
}

type MessageLink struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
	PicUrl      string `xml:"PicUrl"`
	MsgId       string `xml:"MsgId"`
	AgentID     string `xml:"AgentID"`
}
