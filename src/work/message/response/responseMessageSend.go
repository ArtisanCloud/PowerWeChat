package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseMessageSend struct {
	response.ResponseWork

	InvalidUser  string `json:"invaliduser"`  // "userid1|userid2",
	InvalidParty string `json:"invalidparty"` // "partyid1|partyid2",
	InvalidTag   string `json:"invalidtag"`   // "tagid1|tagid2"

	MsgID string `json:"msgid"` // 消息id，用于撤回应用消息

	// 仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回.
	// 应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
	ResponseCode string `json:"response_code"`
}
