package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

// https://developer.work.weixin.qq.com/document/path/90967

const (
	CALLBACK_EVENT_SUBSCRIBE                = "subscribe"
	CALLBACK_EVENT_ENTER_AGENT              = "enter_agent"
	CALLBACK_EVENT_LOCATION                 = "LOCATION"
	CALLBACK_EVENT_BATCH_JOB_RESULT         = "batch_job_result"
	CALLBACK_EVENT_CLICK                    = "click"
	CALLBACK_EVENT_VIEW                     = "view"
	CALLBACK_EVENT_SCANCODE_PUSH            = "scancode_push"
	CALLBACK_EVENT_SCANCODE_WAITMSG         = "scancode_waitmsg"
	CALLBACK_EVENT_PIC_SYSPHOTO             = "pic_sysphoto"
	CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM       = "pic_photo_or_album"
	CALLBACK_EVENT_PIC_WEIXIN               = "pic_weixin"
	CALLBACK_EVENT_LOCATION_SELECT          = "location_select"
	CALLBACK_EVENT_OPEN_APPROVAL_CHANGE     = "open_approval_change"
	CALLBACK_EVENT_SHARE_AGENT_CHANGE       = "share_agent_change"
	CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT = "template_card_menu_event"
	CALLBACK_EVENT_SYS_APPROVAL_CHANGE      = "sys_approval_change"
)

type EventSubscribe struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}

type EventEnterAgent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey"`
	AgentID  string `xml:"AgentID"`
}

type EventLocation struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`
	AgentID   string `xml:"AgentID"`
	AppType   string `xml:"AppType"`
}

type BatchJob struct {
	Text    string `xml:",chardata"`
	JobID   string `xml:"JobId"`
	JobType string `xml:"JobType"`
	ErrCode string `xml:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg"`
}

type EventBatchJobResult struct {
	contract.EventInterface
	models.CallbackMessageHeader
	BatchJob *BatchJob `xml:"BatchJob"`
}

type EventClick struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey"`
	AgentID  string `xml:"AgentID"`
}

type EventView struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey"`
	AgentID  string `xml:"AgentID"`
}

type EventScanCodePush struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey"`
	ScanCodeInfo *ScanCodeInfo `xml:"ScanCodeInfo"`
	AgentID      string        `xml:"AgentID"`
}

type ScanCodeInfo struct {
	Text       string `xml:",chardata"`
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

type EventScancodeWaitMsg struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey"`
	ScanCodeInfo *ScanCodeInfo `xml:"ScanCodeInfo"`
	AgentID      string        `xml:"AgentID"`
}

type EventPicSysPhoto struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey"`
	SendPicsInfo *SendPicsInfo `xml:"SendPicsInfo"`
	AgentID      string        `xml:"AgentID"`
}

type EventPicPhotoOrAlbum struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey"`
	SendPicsInfo *SendPicsInfo `xml:"SendPicsInfo"`
	AgentID      string        `xml:"AgentID"`
}

type Item struct {
	Text      string `xml:",chardata"`
	PicMd5Sum string `xml:"PicMd5Sum"`
}

type PicList struct {
	Text string `xml:",chardata"`
	Item *Item  `xml:"item"`
}

type SendPicsInfo struct {
	Text    string   `xml:",chardata"`
	Count   string   `xml:"Count"`
	PicList *PicList `xml:"PicList"`
}

type EventPicWeixin struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string       `xml:"EventKey"`
	SendPicsInfo SendPicsInfo `xml:"SendPicsInfo"`
	AgentID      string       `xml:"AgentID"`
}

type SendLocationInfo struct {
	Text      string `xml:",chardata"`
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
	PoiName   string `xml:"Poiname"`
}

type EventLocationSelect struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey         string            `xml:"EventKey"`
	SendLocationInfo *SendLocationInfo `xml:"SendLocationInfo"`
	AgentID          string            `xml:"AgentID"`
	AppType          string            `xml:"AppType"`
}

// ----------------------------------------------------------------------------

type Applier struct {
	Text   string `xml:",chardata"`
	UserID string `xml:"UserId"`
	Party  string `xml:"Party"`
}

type Approver struct {
	Text   string `xml:",chardata"`
	UserID string `xml:"UserId"`
}

type Detail struct {
	Text     string    `xml:",chardata"`
	Approver *Approver `xml:"Approver"`
	Speech   string    `xml:"Speech"`
	SpStatus string    `xml:"SpStatus"`
	SpTime   string    `xml:"SpTime"`
}

type SPRecord struct {
	Text         string   `xml:",chardata"`
	SpStatus     string   `xml:"SpStatus"`
	ApproverAttr string   `xml:"ApproverAttr"`
	Details      []Detail `xml:"Details"`
}

type Notifier struct {
	Text   string `xml:",chardata"`
	UserID string `xml:"UserId"`
}

type CommentUserInfo struct {
	Text   string `xml:",chardata"`
	UserID string `xml:"UserId"`
}

type Comments struct {
	Text            string           `xml:",chardata"`
	CommentUserInfo *CommentUserInfo `xml:"CommentUserInfo"`
	CommentTime     string           `xml:"CommentTime"`
	CommentContent  string           `xml:"CommentContent"`
	CommentID       string           `xml:"CommentId"`
}

type ApprovalInfo struct {
	Text              string      `xml:",chardata"`
	SpNO              string      `xml:"SpNo"`
	SpName            string      `xml:"SpName"`
	SpStatus          string      `xml:"SpStatus"`
	TemplateID        string      `xml:"TemplateId"`
	ApplyTime         string      `xml:"ApplyTime"`
	Applier           *Applier    `xml:"Applyer"`
	SpRecord          []*SPRecord `xml:"SpRecord"`
	Notifier          *Notifier   `xml:"Notifyer"`
	Comments          *Comments   `xml:"Comments"`
	StatusChangeEvent string      `xml:"StatuChangeEvent"`
}

type EventOpenApprovalChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID      string        `xml:"AgentID"`
	ApprovalInfo *ApprovalInfo `xml:"ApprovalInfo"`
}

type EventShareAgentChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}

// ----------------------------------------------------------------------------

type OptionID struct {
	Text     string   `xml:",chardata"`
	OptionID []string `xml:"OptionId"`
}

type SelectItem struct {
	Text        string    `xml:",chardata"`
	QuestionKey string    `xml:"QuestionKey"`
	OptionIDs   *OptionID `xml:"OptionIds"`
}

type SelectItems struct {
	Text         string        `xml:",chardata"`
	SelectedItem []*SelectItem `xml:"SelectedItem"`
}

type EventTemplateCardEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey      string       `xml:"EventKey"`
	TaskID        string       `xml:"TaskId"`
	CardType      string       `xml:"CardType"`
	ResponseCode  string       `xml:"ResponseCode"`
	AgentID       string       `xml:"AgentID"`
	SelectedItems *SelectItems `xml:"SelectedItems"`
}

type EventTemplateCardMenuEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string `xml:"EventKey"`
	TaskID       string `xml:"TaskId"`
	CardType     string `xml:"CardType"`
	ResponseCode string `xml:"ResponseCode"`
	AgentID      string `xml:"AgentID"`
}
