package groupChat

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"
)

type Client struct {
	*kernel.BaseClient
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

// 获取客户群列表
// https://developer.work.weixin.qq.com/document/path/92120
func (comp *Client) List(params *request.RequestGroupChatList) (*response.ResponseGroupChatList, error) {

	result := &response.ResponseGroupChatList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/list", params, nil, nil, result)

	return result, err
}

// 获取客户群详情
// https://developer.work.weixin.qq.com/document/path/92122
func (comp *Client) Get(chatID string, needName int) (*response.ResponseGroupChatGet, error) {

	result := &response.ResponseGroupChatGet{}

	options := &object.HashMap{
		"chat_id":   chatID,
		"need_name": needName,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/get?", options, nil, nil, result)

	return result, err
}

// 客户群opengid转换
// https://developer.work.weixin.qq.com/document/path/94822
func (comp *Client) OpenGIDToChatID(openGID string) (*response.ResponseGroupChatOpenGIDToChatID, error) {

	result := &response.ResponseGroupChatOpenGIDToChatID{}

	options := &object.HashMap{
		"opengid": openGID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/opengid_to_chatid?", options, nil, nil, result)

	return result, err
}
