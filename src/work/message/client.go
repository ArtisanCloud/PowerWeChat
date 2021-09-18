package message

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/message/response"
)

type Client struct {
	*kernel.BaseClient
}

func (comp *Client) Message(message *messages.Message) *Messager {

	m := &Messager{
		Client: comp,
	}
	m.Message(message)

	return m
}

// 发送应用消息
// https://work.weixin.qq.com/api/doc/90000/90135/90236
func (comp *Client) Send(messages *power.HashMap) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

	if (*messages)["agentid"] == nil || (*messages)["agentid"].(int) <= 0 {
		config := (*comp.App).GetConfig()
		(*messages)["agentid"] = config.GetInt("agent_id", 0)
	}
	_, err := comp.HttpPostJson("cgi-bin/message/send", messages, nil, nil, result)

	return result, err
}

// 更新模版卡片消息
// https://work.weixin.qq.com/api/doc/90000/90135/94888
func (comp *Client) UpdateTemplateCard(card *power.HashMap) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

	if (*card)["agentid"] == nil || (*card)["agentid"].(int) <= 0 {
		config := (*comp.App).GetConfig()
		(*card)["agentid"] = config.GetInt("agent_id", 0)
	}
	_, err := comp.HttpPostJson("cgi-bin/message/update_template_card", card, nil, nil, result)

	return result, err
}

// Recall 撤回应用消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/94867
func (comp *Client) Recall(msgID string) (*response2.ResponseWork, error) {
	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/message/recall", power.StringMap{
		"msgid": msgID,
	}, nil, nil, result)

	return result, err
}
