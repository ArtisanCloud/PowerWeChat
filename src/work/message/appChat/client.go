package appChat

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/message/appChat/request"
	"github.com/ArtisanCloud/power-wechat/src/work/message/appChat/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 创建群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90245
func (comp *Client) Create(options *request.RequestAppChatCreate) (*response.ResponseAppChatCreate, error) {

	result := &response.ResponseAppChatCreate{}

	_, err := comp.HttpPostJson("cgi-bin/appchat/create", options, nil, nil, result)

	return result, err
}

// 修改群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90246
func (comp *Client) Update(options *request.RequestAppChatUpdate) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/appchat/update", options, nil, nil, result)

	return result, err
}

// 获取群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90247
func (comp *Client) Get(chatID string) (*response.ResponseAppChatGet, error) {

	result := &response.ResponseAppChatGet{}

	params := &object.StringMap{
		"chatid": chatID,
	}

	_, err := comp.HttpGet("cgi-bin/appchat/get", params, nil, result)

	return result, err
}

// 应用推送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90248
func (comp *Client) Send(messages *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/appchat/send", messages, nil, nil, result)

	return result, err
}
