package groupRobot

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/groupRobot/response"
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
func (comp *Client) Send(key string, message *object.HashMap) (*response.ResponseGroupRobotSend, error) {

	result := &response.ResponseGroupRobotSend{}

	comp.Token = nil

	_, err := comp.HttpPostJson("cgi-bin/webhook/send", message, &object.StringMap{
		"key": key,
	}, nil, result)

	return result, err
}
