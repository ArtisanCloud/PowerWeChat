package models

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

type MsgAgentID struct {
	MsgID   uint64 `xml:"MsgId" json:"MsgId"`
	AgentID uint32 `xml:"AgentID" json:"AgentID"`
}

type MessageText struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Content string `xml:"Content" json:"Content"`
	MsgAgentID
}

type MessageImage struct {
	contract.EventInterface
	models.CallbackMessageHeader
	PicUrl  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`
	MsgAgentID
}

type MessageVoice struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaID string `xml:"MediaId"`
	Format  string `xml:"Format"`
	MsgAgentID
}

type MessageLocation struct {
	contract.EventInterface
	models.CallbackMessageHeader
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
	AppType   string `xml:"AppType"`
	MsgAgentID
}

type MessageVideo struct {
	contract.EventInterface
	models.CallbackMessageHeader
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
	MsgAgentID
}

type MessageLink struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
	PicUrl      string `xml:"PicUrl"`
	MsgAgentID
}

// CommonPushData 推送数据通用部分
type CommonPushData struct {
	XMLName      xml.Name `json:"-" xml:"xml"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`           // 消息类型，为固定值 "event"
	Event        string   `json:"Event" xml:"Event"`               // 事件类型
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`     // 小程序的原始 ID
	FromUserName string   `json:"FromUserName" xml:"FromUserName"` // 发送方账号（一个 OpenID，此时发送方是系统账号）
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间（整型），时间戳
}

// MediaCheckAsyncData 媒体内容安全异步审查结果通知
type MediaCheckAsyncData struct {
	CommonPushData
	Appid   string                `json:"appid" xml:"appid"`
	TraceID string                `json:"trace_id" xml:"trace_id"`
	Version int                   `json:"version" xml:"version"`
	Detail  []*MediaCheckDetail   `json:"detail" xml:"detail"`
	ErrCode int                   `json:"errCode" xml:"errCode"`
	ErrMsg  string                `json:"errMsg" xml:"errMsg"`
	Result  MediaCheckAsyncResult `json:"result" xml:"result"`
}

// MediaCheckDetail 检测结果详情
type MediaCheckDetail struct {
	Strategy string `json:"strategy" xml:"strategy"`
	ErrCode  int    `json:"errCode" xml:"errCode"`
	Suggest  string `json:"suggest" xml:"suggest"`
	Label    int    `json:"label" xml:"label"`
	Prob     int    `json:"prob" xml:"prob"`
}

// MediaCheckAsyncResult 检测结果
type MediaCheckAsyncResult struct {
	Suggest string `json:"suggest" xml:"suggest"`
	Label   int    `json:"label" xml:"label"`
}
