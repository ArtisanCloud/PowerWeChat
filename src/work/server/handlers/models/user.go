package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

// https://developer.work.weixin.qq.com/document/path/90970

const (
	CALLBACK_EVENT_CHANGE_CONTACT = "change_contact"

	CALLBACK_EVENT_CHANGE_TYPE_CREATE_USER = "create_user"
	CALLBACK_EVENT_CHANGE_TYPE_UPDATE_USER = "update_user"
	CALLBACK_EVENT_CHANGE_TYPE_DELETE_USER = "delete_user"
)

type EventUserCreate struct {
	contract.EventInterface
	models.CallbackMessageHeader
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
	contract.EventInterface
	models.CallbackMessageHeader
	ChangeType     string `xml:"ChangeType"`
	UserID         string `xml:"UserID"`
	NewUserID      string `xml:"NewUserID"`
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

type EventUserDelete struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID string `xml:"UserID"`
}
