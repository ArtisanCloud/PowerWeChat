package groupRobot

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/groupRobot/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/groupRobot/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

func (comp *Client) Message(message interface{}) (*Messager, error) {
	return NewMessager(comp).Message(message)
}

// https://open.work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B
func (comp *Client) Send(key string, message *power.HashMap) (*response.ResponseGroupRobotSend, error) {

	result := &response.ResponseGroupRobotSend{}

	comp.Token = nil

	_, err := comp.HttpPostJson("cgi-bin/webhook/send", message.ToHashMap(), &object.StringMap{
		"key": key,
	}, nil, result)

	return result, err
}

func (comp *Client) SendText(key string, message *request.GroupRobotMsgText) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "text",
		"text":    message,
	}
	return comp.Send(key, options)
}
func (comp *Client) SendMarkdown(key string, message *request.GroupRobotMsgMarkdown) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype":  "markdown",
		"markdown": message,
	}
	return comp.Send(key, options)
}
func (comp *Client) SendImage(key string, message *request.GroupRobotMsgImage) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "image",
		"image":   message,
	}
	return comp.Send(key, options)
}

// SendNewsArticles 图文类型
func (comp *Client) SendNewsArticles(key string, message []*request.GroupRobotMsgNewsArticles) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "news",
		"news": power.HashMap{
			"articles": message,
		},
	}
	return comp.Send(key, options)
}
func (comp *Client) SendFile(key string, message *request.GroupRobotMsgFile) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype": "file",
		"file":    message,
	}
	return comp.Send(key, options)
}

func (comp *Client) SendTemplateCard(key string, message *request.GroupRobotMsgTemplateCard) (*response.ResponseGroupRobotSend, error) {
	options := &power.HashMap{
		"msgtype":       "template_card",
		"template_card": message,
	}
	return comp.Send(key, options)
}
