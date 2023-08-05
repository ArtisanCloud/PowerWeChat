package models

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

// https://developer.work.weixin.qq.com/document/path/92130

const (
	CALLBACK_EVENT_CHANGE_EXTERNAL_CONTACT               = "change_external_contact"   // 变更外部联系人事件
	CALLBACK_EVENT_CHANGE_TYPE_ADD_EXTERNAL_CONTACT      = "add_external_contact"      // 新增外部联系人事件
	CALLBACK_EVENT_CHANGE_TYPE_EDIT_EXTERNAL_CONTACT     = "edit_external_contact"     // 编辑外部联系人事件
	CALLBACK_EVENT_CHANGE_TYPE_ADD_HALF_EXTERNAL_CONTACT = "add_half_external_contact" // 外部联系人免验证添加成员事件
	CALLBACK_EVENT_CHANGE_TYPE_DEL_EXTERNAL_CONTACT      = "del_external_contact"      // 删除企业客户事件
	CALLBACK_EVENT_CHANGE_TYPE_DEL_FOLLOW_USER           = "del_follow_user"           // 删除跟进成员事件
	CALLBACK_EVENT_CHANGE_TYPE_TRANSFER_FAIL             = "transfer_fail"             // 客户接替失败事件

	CALLBACK_EVENT_CHANGE_EXTERNAL_CHAT        = "change_external_chat" // 客户群创建事件
	CALLBACK_EVENT_UPDATE_DETAIL_ADD_MEMBER    = "add_member"           // 成员入群事件
	CALLBACK_EVENT_UPDATE_DETAIL_DEL_MEMBER    = "del_member"           // 成员退群事件
	CALLBACK_EVENT_UPDATE_DETAIL_CHANGE_OWNER  = "change_owner"         // 群主变更事件
	CALLBACK_EVENT_UPDATE_DETAIL_CHANGE_NAME   = "change_name"          // 群名变更事件
	CALLBACK_EVENT_UPDATE_DETAIL_CHANGE_NOTICE = "change_notice"        // 群公告变更事件

	CALLBACK_EVENT_CHANGE_EXTERNAL_TAG = "change_external_tag" // 企业客户标签变更事件
	CALLBACK_EVENT_CHANGE_TYPE_CREATE  = "create"              // 创建标签事件
	CALLBACK_EVENT_CHANGE_TYPE_UPDATE  = "update"              // 更新标签名称事件
	CALLBACK_EVENT_CHANGE_TYPE_DELETE  = "delete"              // 删除标签事件
	CALLBACK_EVENT_CHANGE_TYPE_DISMISS = "dismiss"             // 解散群事件
	CALLBACK_EVENT_CHANGE_TYPE_SHUFFLE = "shuffle"             // 标签重排事件
	CALLBACK_EVENT_TAG_TYPE            = "tag"                 // 标签事件
	CALLBACK_EVENT_TAG_TYPE_GROUP      = "tag_group"           // 标签组事件
)

type EventExternalUserAdd struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
	WelcomeCode    string `xml:"WelcomeCode"`
}

type EventExternalUserEdit struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}

type EventExternalUserAddHalf struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
	WelcomeCode    string `xml:"WelcomeCode"`
}

type EventExternalUserDel struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	Source         string `xml:"Source"`
}

type EventExternalUserDelFollowUser struct {
	contract.EventInterface
	models.CallbackMessageHeader
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}

type EventExternalChatCreate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatID string `xml:"ChatId"`
}

type EventExternalChatUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatID       string `xml:"ChatId"`
	UpdateDetail string `xml:"UpdateDetail"`
	JoinScene    string `xml:"JoinScene"`
	QuitScene    string `xml:"QuitScene"`
	MemChangeCnt string `xml:"MemChangeCnt"`
}

type EventExternalChatDismiss struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatID string `xml:"ChatId"`
}

// Deprecated: Use EventExternalChatUpdate instead.
type EventExternalUserUpdateAddMember struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatID       string `xml:"ChatId"`
	ChangeType   string `xml:"ChangeType"`
	UpdateDetail string `xml:"UpdateDetail"`
	JoinScene    string `xml:"JoinScene"`
	QuitScene    string `xml:"QuitScene"`
	MemChangeCnt string `xml:"MemChangeCnt"`
}

// Deprecated: Use EventExternalChatDismiss instead.
type EventExternalUserDismiss struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ChatID string `xml:"ChatId"`
}

type EventExternalTransferFail struct {
	contract.EventInterface
	models.CallbackMessageHeader
	FailReason     string `xml:"FailReason"`
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}

type EventExternalUserTagCreate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID         string `xml:"Id"`
	TagType    string `xml:"TagType"`
	StrategyID string `xml:"StrategyId"`
}

type EventExternalUserTagUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID         string `xml:"Id"`
	TagType    string `xml:"TagType"`
	StrategyID string `xml:"StrategyId"`
}

type EventExternalUserTagDelete struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID         string `xml:"Id"`
	TagType    string `xml:"TagType"`
	StrategyID string `xml:"StrategyId"`
}

type EventExternalUserTagShuffle struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID         string `xml:"Id"`
	StrategyID string `xml:"StrategyId"`
	ChangeType string `xml:"ChangeType"`
}
