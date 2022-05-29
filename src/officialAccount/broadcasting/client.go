package broadcasting

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/broadcasting/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/broadcasting/response"
)

const BROADCASTING_PREVIEW_BY_OPENID string = "touser"
const BROADCASTING_PREVIEW_BY_NAME string = "towxname"

type Client struct {
	*kernel.BaseClient
}

// 根据标签进行群发 【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Send(message *power.HashMap) (*response.ResponseBroadcastingSend, error) {

	if (*message)["filter"] == nil && (*message)["touser"] == nil {
		return nil, errors.New("the message reception object is not specified")
	}

	result := &response.ResponseBroadcastingSend{}

	api := "cgi-bin/message/mass/sendall"
	if (*message)["touser"] != nil {
		api = "cgi-bin/message/mass/send"
	}

	_, err := comp.HttpPostJson(api, message, nil, nil, result)
	return result, err
}

// 预览接口【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Preview(data *object.HashMap) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("cgi-bin/message/mass/preview", data, nil, nil, result)
	return result, err
}

// 删除群发【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Delete(msgID string, index int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_id":      msgID,
		"article_idx": index,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/mass/delete", params, nil, nil, result)
	return result, err
}

// 查询群发消息发送状态【订阅号与服务号认证后均可用】
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
func (comp *Client) Status(msgID string) (*response.ResponseBroadcastingStatus, error) {

	result := &response.ResponseBroadcastingStatus{}

	params := &object.HashMap{
		"msg_id": msgID,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/mass/get", params, nil, nil, result)
	return result, err
}

func (comp *Client) SendText(message string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewText(message), reception, attributes)
}

func (comp *Client) SendNews(mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewMedia(mediaID, "mpnews", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendVoice(mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewMedia(mediaID, "voice", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendImage(mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewImage(mediaID, &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendVideo(mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewMedia(mediaID, "mpvideo", &power.HashMap{}), reception, attributes)
}

func (comp *Client) SendCard(mediaID string, reception *request.Reception, attributes *power.HashMap) (interface{}, error) {

	return comp.SendMessage(messages.NewCard(mediaID), reception, attributes)
}

func (comp *Client) PreviewText(message string, reception *request.Reception, method string) (*response2.ResponseOfficialAccount, error) {

	return comp.PreviewMessage(messages.NewText(message), reception, method)
}

func (comp *Client) PreviewMessage(message contract.MessageInterface, reception *request.Reception, method string) (*response2.ResponseOfficialAccount, error) {

	previewMessage, err := NewMessengerBuilder().SetMessage(message).BuildForPreview(method, reception)
	if err != nil {
		return nil, err
	}
	return comp.Preview(previewMessage)
}

func (comp *Client) SendMessage(message contract.MessageInterface, reception *request.Reception, attribute *power.HashMap) (interface{}, error) {

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
	return comp.Send(powerMsg)

}
