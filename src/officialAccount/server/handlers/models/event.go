package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

const (
	CALLBACK_EVENT_SUBSCRIBE                = "subscribe"
	CALLBACK_EVENT_UNSUBSCRIBE              = "unsubscribe"
	CALLBACK_EVENT_ENTER_AGENT              = "enter_agent"
	CALLBACK_EVENT_LOCATION                 = "LOCATION"
	CALLBACK_EVENT_BATCH_JOB_RESULT         = "batch_job_result"
	CALLBACK_EVENT_CLICK                    = "CLICK"
	CALLBACK_EVENT_VIEW                     = "VIEW"
	CALLBACK_EVENT_SCAN                     = "SCAN"
	CALLBACK_EVENT_SCANCODE_PUSH            = "scancode_push"
	CALLBACK_EVENT_SCANCODE_WAITMSG         = "scancode_waitmsg"
	CALLBACK_EVENT_PIC_SYSPHOTO             = "pic_sysphoto"
	CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM       = "pic_photo_or_album"
	CALLBACK_EVENT_PIC_WEIXIN               = "pic_weixin"
	CALLBACK_EVENT_LOCATION_SELECT          = "location_select"
	CALLBACK_EVENT_OPEN_APPROVAL_CHANGE     = "open_approval_change"
	CALLBACK_EVENT_SHARE_AGENT_CHANGE       = "share_agent_change"
	CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT = "template_card_menu_event"
	CALLBACK_EVENT_TEMPLATE_SEND_JOB_FINISH = "TEMPLATESENDJOBFINISH"
	CALLBACK_EVENT_VIEW_MINIPROGRAM         = "view_miniprogram"
	CALLBACK_EVENT_MASS_SEND_JOB_FINISH     = "MASSSENDJOBFINISH"
)

type EventSubscribe struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey"`
	AgentID  string `xml:"AgentID"`
	Ticket   string `xml:"Ticket"`
}

type EventUnSubscribe struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
	Ticket  string `xml:"Ticket"`
}

type EventScan struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID  string `xml:"AgentID"`
	Ticket   string `xml:"Ticket"`
	EventKey string `xml:"EventKey"`
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

type EventBatchJobResult struct {
	contract.EventInterface
	models.CallbackMessageHeader
	BatchJob struct {
		Text    string `xml:",chardata"`
		JobID   string `xml:"JobId"`
		JobType string `xml:"JobType"`
		ErrCode string `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
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
	EventKey     string `xml:"EventKey"`
	ScanCodeInfo struct {
		Text       string `xml:",chardata"`
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`
	AgentID string `xml:"AgentID"`
}

type EventScancodeWaitMsg struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string `xml:"EventKey"`
	ScanCodeInfo struct {
		Text       string `xml:",chardata"`
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`
	AgentID string `xml:"AgentID"`
}

type EventPicSysPhoto struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string `xml:"EventKey"`
	SendPicsInfo struct {
		Text    string `xml:",chardata"`
		Count   string `xml:"Count"`
		PicList struct {
			Text string `xml:",chardata"`
			Item struct {
				Text      string `xml:",chardata"`
				PicMd5Sum string `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
	AgentID string `xml:"AgentID"`
}

type EventPicPhotoOrAlbum struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string `xml:"EventKey"`
	SendPicsInfo struct {
		Text    string `xml:",chardata"`
		Count   string `xml:"Count"`
		PicList struct {
			Text string `xml:",chardata"`
			Item struct {
				Text      string `xml:",chardata"`
				PicMd5Sum string `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
	AgentID string `xml:"AgentID"`
}

type EventPicWeixin struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string `xml:"EventKey"`
	SendPicsInfo struct {
		Text    string `xml:",chardata"`
		Count   string `xml:"Count"`
		PicList struct {
			Text string `xml:",chardata"`
			Item struct {
				Text      string `xml:",chardata"`
				PicMd5Sum string `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
	AgentID string `xml:"AgentID"`
}

type EventLocationSelect struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey         string `xml:"EventKey"`
	SendLocationInfo struct {
		Text      string `xml:",chardata"`
		LocationX string `xml:"Location_X"`
		LocationY string `xml:"Location_Y"`
		Scale     string `xml:"Scale"`
		Label     string `xml:"Label"`
		Poiname   string `xml:"Poiname"`
	} `xml:"SendLocationInfo"`
	AgentID string `xml:"AgentID"`
	AppType string `xml:"AppType"`
}

type EventOpenApprovalChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID      string `xml:"AgentID"`
	ApprovalInfo struct {
		Text           string `xml:",chardata"`
		ThirdNo        string `xml:"ThirdNo"`
		OpenSpName     string `xml:"OpenSpName"`
		OpenTemplateID string `xml:"OpenTemplateId"`
		OpenSpStatus   string `xml:"OpenSpStatus"`
		ApplyTime      string `xml:"ApplyTime"`
		ApplyUserName  string `xml:"ApplyUserName"`
		ApplyUserID    string `xml:"ApplyUserId"`
		ApplyUserParty string `xml:"ApplyUserParty"`
		ApplyUserImage string `xml:"ApplyUserImage"`
		ApprovalNodes  struct {
			Text         string `xml:",chardata"`
			ApprovalNode struct {
				Text       string `xml:",chardata"`
				NodeStatus string `xml:"NodeStatus"`
				NodeAttr   string `xml:"NodeAttr"`
				NodeType   string `xml:"NodeType"`
				Items      struct {
					Text string `xml:",chardata"`
					Item struct {
						Text       string `xml:",chardata"`
						ItemName   string `xml:"ItemName"`
						ItemUserID string `xml:"ItemUserId"`
						ItemImage  string `xml:"ItemImage"`
						ItemStatus string `xml:"ItemStatus"`
						ItemSpeech string `xml:"ItemSpeech"`
						ItemOpTime string `xml:"ItemOpTime"`
					} `xml:"Item"`
				} `xml:"Items"`
			} `xml:"ApprovalNode"`
		} `xml:"ApprovalNodes"`
		NotifyNodes struct {
			Text       string `xml:",chardata"`
			NotifyNode struct {
				Text       string `xml:",chardata"`
				ItemName   string `xml:"ItemName"`
				ItemUserID string `xml:"ItemUserId"`
				ItemImage  string `xml:"ItemImage"`
			} `xml:"NotifyNode"`
		} `xml:"NotifyNodes"`
		Approverstep string `xml:"approverstep"`
	} `xml:"ApprovalInfo"`
}

type EventShareAgentChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}

type EventTemplateCardEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey      string `xml:"EventKey"`
	TaskID        string `xml:"TaskId"`
	CardType      string `xml:"CardType"`
	ResponseCode  string `xml:"ResponseCode"`
	AgentID       string `xml:"AgentID"`
	SelectedItems struct {
		Text         string `xml:",chardata"`
		SelectedItem []struct {
			Text        string `xml:",chardata"`
			QuestionKey string `xml:"QuestionKey"`
			OptionIds   struct {
				Text     string   `xml:",chardata"`
				OptionID []string `xml:"OptionId"`
			} `xml:"OptionIds"`
		} `xml:"SelectedItem"`
	} `xml:"SelectedItems"`
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
