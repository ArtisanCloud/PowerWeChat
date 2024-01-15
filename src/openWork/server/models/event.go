package models

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	kernelModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

type InfoType = string

const (
	InfoTypeSuiteTicket        InfoType = "suite_ticket"
	InfoTypeCreateAuth         InfoType = "create_auth"
	InfoTypeChangeAuth         InfoType = "change_auth"
	InfoTypeCancelAuth         InfoType = "cancel_auth"
	InfoTypeChangeContact      InfoType = "change_contact"
	InfoTypeShareAgentChange   InfoType = "share_agent_change"
	InfoTypeShareChainChange   InfoType = "share_chain_change"
	InfoTypeResetPermanentCode InfoType = "reset_permanent_code"
	InfoTypeCorpArchAuth       InfoType = "corp_arch_auth"
	InfoTypeApproveSpecialAuth InfoType = "approve_special_auth"
	InfoTypeCancelSpecialAuth  InfoType = "cancel_special_auth"
)

type ChangeType = string

const (
	ChangeTypeCreateUser  ChangeType = "create_user"
	ChangeTypeUpdateUser  ChangeType = "update_user"
	ChangeTypeDeleteUser  ChangeType = "delete_user"
	ChangeTypeCreateParty ChangeType = "create_party"
	ChangeTypeUpdateParty ChangeType = "update_party"
	ChangeTypeDeleteParty ChangeType = "delete_party"
	ChangeTypeUpdateTag   ChangeType = "update_tag"
)

func DecodeEvent(bs []byte) (IEvent, error) {
	var msg BaseEvent
	if err := xml.Unmarshal(bs, &msg); err != nil {
		return nil, err
	}
	ev, err := msg.ToEvent()
	if err != nil {
		return msg, fmt.Errorf("%w: %s", err, string(bs))
	}
	if err := xml.Unmarshal(bs, ev); err != nil {
		return msg, fmt.Errorf("%w: %s", err, string(bs))
	}
	return ev, nil
}

type IEvent interface {
	contract.EventInterface
	GetSuiteID() string
	GetInfoType() InfoType
	GetTimestamp() int64
	GetChangeType() ChangeType
}

type BaseEvent struct {
	kernelModels.CallbackMessageHeader
	XMLName    xml.Name   `xml:"xml"`
	SuiteID    string     `xml:"SuiteId"`
	InfoType   InfoType   `xml:"InfoType"`
	Timestamp  int64      `xml:"TimeStamp"`
	ChangeType ChangeType `xml:"ChangeType,omitempty"`
}

func (ev BaseEvent) GetSuiteID() string {
	return ev.SuiteID
}

func (ev BaseEvent) GetInfoType() string {
	return ev.InfoType
}

func (ev BaseEvent) GetChangeType() ChangeType {
	return ev.ChangeType
}

func (ev BaseEvent) GetTimestamp() int64 {
	return ev.Timestamp
}

func (msg BaseEvent) ToEvent() (IEvent, error) {
	switch msg.GetInfoType() {
	case InfoTypeSuiteTicket:
		return new(EventSuiteTicket), nil
	case InfoTypeCreateAuth:
		return new(EventCreateAuth), nil
	case InfoTypeChangeAuth:
		return new(EventChangeAuth), nil
	case InfoTypeCancelAuth:
		return new(EventCancelAuth), nil
	case InfoTypeChangeContact:
		switch msg.GetChangeType() {
		case ChangeTypeCreateUser, ChangeTypeUpdateUser, ChangeTypeDeleteUser:
			return new(EventUser), nil
		case ChangeTypeCreateParty, ChangeTypeUpdateParty, ChangeTypeDeleteParty:
			return new(EventParty), nil
		case ChangeTypeUpdateTag:
			return new(EventUpdateTag), nil
		default:
			return nil, errors.New("unknown event")
		}
	case InfoTypeShareAgentChange, InfoTypeShareChainChange:
		return new(EventShareChange), nil
	case InfoTypeResetPermanentCode:
		return new(EventResetPermanentCode), nil
	case InfoTypeCorpArchAuth:
		return new(EventCorpArchAuth), nil
	case InfoTypeApproveSpecialAuth, InfoTypeCancelSpecialAuth:
		return new(EventSpecialAuth), nil
	default:
		return nil, errors.New("unknown event")
	}
}

type EventSuiteTicket struct {
	BaseEvent
	SuiteTicket string `xml:"SuiteTicket"`
}

type EventCreateAuth struct {
	BaseEvent
	AuthCode string `xml:"AuthCode"`
	State    string `xml:"State"`
}

type EventChangeAuth struct {
	BaseEvent
	AuthCorpID string `xml:"AuthCorpId"`
	State      string `xml:"State"`
}

type EventCancelAuth struct {
	BaseEvent
	AuthCorpID string `xml:"AuthCorpId"`
}

type EventUser struct {
	BaseEvent
	AuthCorpID     string `xml:"AuthCorpId"`
	UserID         string `xml:"UserID"`
	OpenUserID     string `xml:"OpenUserID"`
	Name           string `xml:"Name"`
	Department     string `xml:"Department"`
	MainDepartment uint64 `xml:"MainDepartment"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	DirectLeader   string `xml:"DirectLeader"`
	Mobile         string `xml:"Mobile"`
	Position       string `xml:"Position"`
	Gender         string `xml:"Gender"`
	Email          string `xml:"Email"`
	BizMail        string `xml:"BizMail"`
	Avatar         string `xml:"Avatar"`
	Alias          string `xml:"Alias"`
	Telephone      string `xml:"Telephone"`
	ExtAttr        string `xml:"ExtAttr>Item"`
}

type ExtAttr struct {
	Name string `xml:"Name"`
	Type string `xml:"Type"`
	Text string `xml:"Text>Value"`
	Web  *Web   `xml:"Web"`
}

type Web struct {
	Title string `xml:"Title"`
	Url   string `xml:"Url"`
}

type EventParty struct {
	BaseEvent
	AuthCorpID string `xml:"AuthCorpId"`
	ID         uint64 `xml:"Id"`
	Name       string `xml:"Name"`
	ParentID   uint64 `xml:"ParentId"`
	OrderID    int    `xml:"OrderId"`
}

type EventUpdateTag struct {
	BaseEvent
	AuthCorpID    string `xml:"AuthCorpId"`
	TagID         uint64 `xml:"TagId"`
	AddUserItems  string `xml:"AddUserItems"`
	DelUserItems  string `xml:"DelUserItems"`
	AddPartyItems string `xml:"AddPartyItems"`
	DelPartyItems string `xml:"DelPartyItems"`
}

type EventShareChange struct {
	BaseEvent
	AppID   uint64 `xml:"AppId"`
	CropID  string `xml:"CorpId"`
	AgentID uint64 `xml:"AgentId"`
}

type EventResetPermanentCode struct {
	BaseEvent
	AuthCode string `xml:"AuthCode"`
}

type EventCorpArchAuth struct {
	BaseEvent
	AuthCorpID string `xml:"AuthCorpId"`
}

type EventSpecialAuth struct {
	BaseEvent
	AuthCorpID string `xml:"AuthCorpId"`
	AuthType   string `xml:"AuthType"`
}
