package groupChat

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/groupChat/request"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/groupChat/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取客户群列表
// https://work.weixin.qq.com/api/doc/90000/90135/92120
func (comp  *Client) List(params *request.RequestGroupChatList) (*response.ResponseGroupChatList, error) {

	result := &response.ResponseGroupChatList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/list", params, nil, nil, result)

	return result, err
}

// 获取客户群详情
// https://work.weixin.qq.com/api/doc/90000/90135/92122
func (comp  *Client) Get(chatID string, needName bool) (*response.ResponseGroupChatGet, error) {

	result := &response.ResponseGroupChatGet{}

	options := &object.HashMap{
		"chat_id":   chatID,
		"need_name": needName,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/get?", options, nil, nil, result)

	return result, err
}

// 客户群opengid转换
// https://work.weixin.qq.com/api/doc/90000/90135/94822
func (comp  *Client) OpenGIDToChatID(openGID string) (*response.ResponseGroupChatOpenGIDToChatID, error) {

	result := &response.ResponseGroupChatOpenGIDToChatID{}

	options := &object.HashMap{
		"opengid": openGID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/opengid_to_chatid?", options, nil, nil, result)

	return result, err
}
