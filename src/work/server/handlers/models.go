package handlers

import "encoding/xml"

type CallbackInterface interface {
	GetMsgType() string
	GetEvent() string
}

type CallbackHeader struct {
	CallbackInterface
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
}

func (ch CallbackHeader) GetMsgType() string {
	return ch.MsgType
}

func (ch CallbackHeader) GetEvent() string {
	return ch.Event
}

type EventUserCreate struct {
	CallbackInterface
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	CallbackHeader
	ChangeType     string `xml:"ChangeType"`
	UserID         string `xml:"UserID"`
	Name           string `xml:"Name"`
	Department     string `xml:"Department"`
	MainDepartment string `xml:"MainDepartment"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	Position       string `xml:"Position"`
	Mobile         string `xml:"Mobile"`
	Gender         string `xml:"Gender"`
	Email          string `xml:"Email"`
	Status         string `xml:"Status"`
	Avatar         string `xml:"Avatar"`
	Alias          string `xml:"Alias"`
	Telephone      string `xml:"Telephone"`
	Address        string `xml:"Address"`
	ExtAttr        struct {
		Text string `xml:",chardata"`
		Item []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"Name"`
			Type     string `xml:"Type"`
			Text     struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value"`
			} `xml:"Text"`
			Web struct {
				Text  string `xml:",chardata"`
				Title string `xml:"Title"`
				URL   string `xml:"Url"`
			} `xml:"Web"`
		} `xml:"Item"`
	} `xml:"ExtAttr"`
}

type EventUserUpdate struct {
	CallbackInterface
	XMLName        xml.Name `xml:"xml"`
	Text           string   `xml:",chardata"`
	CallbackHeader
	ChangeType     string   `xml:"ChangeType"`
	UserID         string   `xml:"UserID"`
	NewUserID      string   `xml:"NewUserID"`
	Name           string   `xml:"Name"`
	Department     string   `xml:"Department"`
	MainDepartment string   `xml:"MainDepartment"`
	IsLeaderInDept string   `xml:"IsLeaderInDept"`
	Position       string   `xml:"Position"`
	Mobile         string   `xml:"Mobile"`
	Gender         string   `xml:"Gender"`
	Email          string   `xml:"Email"`
	Status         string   `xml:"Status"`
	Avatar         string   `xml:"Avatar"`
	Alias          string   `xml:"Alias"`
	Telephone      string   `xml:"Telephone"`
	Address        string   `xml:"Address"`
	ExtAttr        struct {
		Text string `xml:",chardata"`
		Item []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"Name"`
			Type     string `xml:"Type"`
			Text     struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value"`
			} `xml:"Text"`
			Web struct {
				Text  string `xml:",chardata"`
				Title string `xml:"Title"`
				URL   string `xml:"Url"`
			} `xml:"Web"`
		} `xml:"Item"`
	} `xml:"ExtAttr"`
}
