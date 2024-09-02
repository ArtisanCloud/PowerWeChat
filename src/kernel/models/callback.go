package models

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
)

type Callback struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName string   `xml:"ToUserName" json:"ToUserName"`
	AgentID    string   `xml:"AgentID" json:"AgentID"`
	Encrypt    string   `xml:"Encrypt" json:"Encrypt"`
}

const (
	CALLBACK_MSG_TYPE_EVENT    = "event"
	CALLBACK_MSG_TYPE_TEXT     = "text"
	CALLBACK_MSG_TYPE_IMAGE    = "image"
	CALLBACK_MSG_TYPE_VOICE    = "voice"
	CALLBACK_MSG_TYPE_VIDEO    = "video"
	CALLBACK_MSG_TYPE_LOCATION = "location"
	CALLBACK_MSG_TYPE_LINK     = "link"
)

type CallbackMessageHeader struct {
	contract.EventInterface
	XMLName      xml.Name `xml:"xml"`
	Text         string   `xml:",chardata"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	FromUserName string   `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int      `xml:"CreateTime" json:"CreateTime"`
	MsgType      string   `xml:"MsgType" json:"MsgType"`
	Event        string   `xml:"Event" json:"Event"`
	ChangeType   string   `xml:"ChangeType" json:"ChangeType"`
	EventKey     string   `xml:"EventKey,omitempty" json:"EventKey,omitempty"`
	Content      []byte
}

func (header CallbackMessageHeader) GetToUserName() string {
	return header.ToUserName
}

func (header CallbackMessageHeader) GetFromUserName() string {
	return header.FromUserName
}

func (header CallbackMessageHeader) GetCreateTime() int {
	return header.CreateTime
}

func (header CallbackMessageHeader) GetMsgType() string {
	return header.MsgType
}

func (header CallbackMessageHeader) GetEvent() string {
	return header.Event
}

func (header CallbackMessageHeader) GetChangeType() string {
	return header.ChangeType
}

func (header CallbackMessageHeader) ReadMessage(msg interface{}) error {
	return xml.Unmarshal(header.Content, msg)
}

func (header CallbackMessageHeader) GetContent() []byte {
	return header.Content
}
