package message

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

func (comp *Client) Message(message *messages.Message) *Messager {

	m := &Messager{
		Client: comp,
	}
	m.Message(message)

	return m
}

// 发送应用消息
// https://developer.work.weixin.qq.com/document/path/90236
func (comp *Client) Send(ctx context.Context, messages interface{}) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/send", messages, nil, nil, result)

	return result, err
}

// SendText 文本消息
func (comp *Client) SendText(ctx context.Context, messages *request.RequestMessageSendText) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendImage 图片消息
func (comp *Client) SendImage(ctx context.Context, messages *request.RequestMessageSendImage) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendVoice 语音消息
func (comp *Client) SendVoice(ctx context.Context, messages *request.RequestMessageSendVoice) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendVideo 视频消息
func (comp *Client) SendVideo(ctx context.Context, messages *request.RequestMessageSendVideo) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendFile 文件消息
func (comp *Client) SendFile(ctx context.Context, messages *request.RequestMessageSendFile) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendTextCard 文本卡片消息
func (comp *Client) SendTextCard(ctx context.Context, messages *request.RequestMessageSendTextCard) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendNews 图文消息
func (comp *Client) SendNews(ctx context.Context, messages *request.RequestMessageSendNews) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendMpNews 图文消息（mpnews）
func (comp *Client) SendMpNews(ctx context.Context, messages *request.RequestMessageSendMPNews) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendMarkdown markdown消息
func (comp *Client) SendMarkdown(ctx context.Context, messages *request.RequestMessageSendMarkdown) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendMiniProgramNotice 小程序通知消息
func (comp *Client) SendMiniProgramNotice(ctx context.Context, messages *request.RequestMessageSendMiniProgramNotice) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// SendTemplateCard 发送卡片模版
func (comp *Client) SendTemplateCard(ctx context.Context, messages *request.RequestMessageSendTemplateCard) (*response.ResponseMessageSend, error) {
	return comp.Send(ctx, messages)
}

// 更新模版卡片消息
// https://developer.work.weixin.qq.com/document/path/94888
func (comp *Client) UpdateTemplateCard(ctx context.Context, card *power.HashMap) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

	if (*card)["agentid"] == nil || (*card)["agentid"].(int) <= 0 {
		config := (*comp.BaseClient.App).GetConfig()
		(*card)["agentid"] = config.GetInt("agent_id", 0)
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/update_template_card", card, nil, nil, result)

	return result, err
}

// Recall 撤回应用消息
// https://developer.work.weixin.qq.com/document/path/94867
func (comp *Client) Recall(ctx context.Context, msgID string) (*response2.ResponseWork, error) {
	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/recall", power.StringMap{
		"msgid": msgID,
	}, nil, nil, result)

	return result, err
}
