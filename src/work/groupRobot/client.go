package groupRobot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/response"
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

func (comp *Client) Message(message interface{}) (*Messager, error) {
	return NewMessager(comp).Message(message)
}

// 群机器人配置说明

// https://developer.work.weixin.qq.com/document/path/91770#文本类型
func (comp *Client) UploadMedia(ctx context.Context, key string, path string) (*response.ResponseGroupRobotUploadMedia, error) {

	result := &response.ResponseGroupRobotUploadMedia{}

	comp.BaseClient.Token = nil

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	} else {
		return nil, errors.New("path is empty")
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "cgi-bin/webhook/upload_media", files, nil, &object.StringMap{
		"key":  key,
		"type": "file",
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/91770#文本类型
func (comp *Client) Send(ctx context.Context, key string, message *power.HashMap) (*response.ResponseGroupRobotSend, error) {

	result := &response.ResponseGroupRobotSend{}

	comp.BaseClient.Token = nil

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/webhook/send", message.ToHashMap(), &object.StringMap{
		"key": key,
	}, nil, result)

	return result, err
}

func (comp *Client) SendText(ctx context.Context, key string, message *request.GroupRobotMsgText) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "text",
		"text":    message,
	}
	return comp.Send(ctx, key, options)
}
func (comp *Client) SendMarkdown(ctx context.Context, key string, message *request.GroupRobotMsgMarkdown) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype":  "markdown",
		"markdown": message,
	}
	return comp.Send(ctx, key, options)
}
func (comp *Client) SendImage(ctx context.Context, key string, message *request.GroupRobotMsgImage) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "image",
		"image":   message,
	}
	return comp.Send(ctx, key, options)
}

// SendNewsArticles 图文类型
func (comp *Client) SendNewsArticles(ctx context.Context, key string, message []*request.GroupRobotMsgNewsArticles) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "news",
		"news": power.HashMap{
			"articles": message,
		},
	}
	return comp.Send(ctx, key, options)
}
func (comp *Client) SendFile(ctx context.Context, key string, message *request.GroupRobotMsgFile) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "file",
		"file":    message,
	}
	return comp.Send(ctx, key, options)
}

func (comp *Client) SendTemplateCard(ctx context.Context, key string, message *request.GroupRobotMsgTemplateCard) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype":       "template_card",
		"template_card": message,
	}
	return comp.Send(ctx, key, options)
}
