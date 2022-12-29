package broadcasting

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/broadcasting/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/broadcasting/response"
)

const BROADCASTING_PREVIEW_BY_OPENID string = "touser"
const BROADCASTING_PREVIEW_BY_NAME string = "towxname"

type Client struct {
	BaseClient *kernel.BaseClient
}

// 根据标签进行群发 【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Send(ctx *context.Context, message *power.HashMap) (*response.ResponseBroadcastingSend, error) {

	if (*message)["filter"] == nil && (*message)["touser"] == nil {
		return nil, errors.New("the message reception object is not specified")
	}

	result := &response.ResponseBroadcastingSend{}

	api := "cgi-bin/message/mass/sendall"
	if (*message)["touser"] != nil {
		api = "cgi-bin/message/mass/send"
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, api, message, nil, nil, result)
	return result, err
}

// 预览接口【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Preview(ctx *context.Context, data *object.HashMap) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/mass/preview", data, nil, nil, result)
	return result, err
}

// 删除群发【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Delete(ctx *context.Context, msgID string, index int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_id":      msgID,
		"article_idx": index,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/mass/delete", params, nil, nil, result)
	return result, err
}

// 查询群发消息发送状态【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Status(ctx *context.Context, msgID string) (*response.ResponseBroadcastingStatus, error) {

	result := &response.ResponseBroadcastingStatus{}

	params := &object.HashMap{
		"msg_id": msgID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/mass/get", params, nil, nil, result)
	return result, err
}

func (comp *Client) SendText(ctx *context.Context, message string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewText(message), reception, attributes)
}

func (comp *Client) SendNews(ctx *context.Context, mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewMedia(mediaID, "mpnews", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendVoice(ctx *context.Context, mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewMedia(mediaID, "voice", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendImage(ctx *context.Context, mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewImage(mediaID, &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendVideo(ctx *context.Context, mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewMedia(mediaID, "mpvideo", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendCard(ctx *context.Context, mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(ctx, messages.NewCard(mediaID), reception, attributes)
}

func (comp *Client) PreviewText(ctx *context.Context, message string, reception *request.Reception, method string) (*response2.ResponseOfficialAccount, error) {

	return comp.PreviewMessage(ctx, messages.NewText(message), reception, method)
}

func (comp *Client) PreviewMessage(ctx *context.Context, message contract.MessageInterface, reception *request.Reception, method string) (*response2.ResponseOfficialAccount, error) {

	previewMessage, err := NewMessengerBuilder().SetMessage(message).BuildForPreview(method, reception)
	if err != nil {
		return nil, err
	}
	return comp.Preview(ctx, previewMessage)
}

func (comp *Client) SendMessage(ctx *context.Context, message contract.MessageInterface, reception *request.Reception, attribute *power.HashMap) (interface{}, error) {

	messageBuilder, err := NewMessengerBuilder().SetMessage(message).With(attribute)
	if err != nil {
		return nil, err
	}
	messageBuilder.ToAll()

	if reception.Filter != nil && reception.Filter.TagID > 0 {
		messageBuilder.ToTag(reception.Filter.TagID)

	} else if len(reception.ToUser) > 0 {
		messageBuilder.ToUsers(reception.ToUser)
	}

	msg, err := messageBuilder.Build(nil)
	if err != nil {
		return nil, err
	}

	powerMsg, _ := power.HashMapToPower(msg)
	return comp.Send(ctx, powerMsg)

}
