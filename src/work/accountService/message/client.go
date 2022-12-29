package message

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/message/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/message/response"
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

// 读取消息
// https://developer.work.weixin.qq.com/document/path/94670
func (comp *Client) SyncMsg(ctx context.Context, cursor string, token string, limit int) (*response.ResponseMessageSyncMsg, error) {

	result := &response.ResponseMessageSyncMsg{}

	options := &object.HashMap{
		"cursor": cursor,
		"token":  token,
		"limit":  limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/sync_msg", options, nil, nil, result)

	return result, err
}

// 发送消息
// https://developer.work.weixin.qq.com/document/path/90236
func (comp *Client) SendMsg(ctx context.Context, messages *request.RequestAccountServiceSendMsg) (*response.ResponseAccountServiceSendMsg, error) {

	result := &response.ResponseAccountServiceSendMsg{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/send_msg", messages, nil, nil, result)

	return result, err
}

// 发送消息
// https://developer.work.weixin.qq.com/document/path/95122
func (comp *Client) SendMsgOnEvent(ctx context.Context, messages *request.RequestAccountServiceSendMsgOnEvent) (*response.ResponseAccountServiceSendMsgOnEvent, error) {

	result := &response.ResponseAccountServiceSendMsgOnEvent{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/send_msg_on_event", messages, nil, nil, result)

	return result, err
}
