package message

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/work/message/request"
	"github.com/ArtisanCloud/PowerWeChat/src/work/message/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
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

func (comp *Client) send(messages interface{}) (*response.ResponseMessageSend, error) {
	return comp.Send(messages.(*power.HashMap))
}

// SendText 文本消息
func (comp *Client) SendText(messages *request.RequestMessageSendText) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}
// SendImage 图片消息
func (comp *Client) SendImage(messages *request.RequestMessageSendImage) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendVoice 语音消息
func (comp *Client) SendVoice(messages *request.RequestMessageSendVoice) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendVideo 视频消息
func (comp *Client) SendVideo(messages *request.RequestMessageSendVideo) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendFile 文件消息
func (comp *Client) SendFile(messages *request.RequestMessageSendFile) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendTextCard 文本卡片消息
func (comp *Client) SendTextCard(messages *request.RequestMessageSendTextCard) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendNews 图文消息
func (comp *Client) SendNews(messages *request.RequestMessageSendNews) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendMpNews 图文消息（mpnews）
func (comp *Client) SendMpNews(messages *request.RequestMessageSendMPNews) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendMarkdown markdown消息
func (comp *Client) SendMarkdown(messages *request.RequestMessageSendMarkdown) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendMiniProgramNotice 小程序通知消息
func (comp *Client) SendMiniProgramNotice(messages *request.RequestMessageSendMiniProgramNotice) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
}

// SendTemplateCard 发送卡片模版
func (comp *Client) SendTemplateCard(messages *request.RequestMessageSendTemplateCard) (*response.ResponseMessageSend, error) {
	return comp.send(messages)
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
