package groupChat

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
)

type GroupChat struct {
	*kernel.BaseClient
}

func NewGroupChat(app kernel.ApplicationInterface) *GroupChat {
	return &GroupChat{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取客户群列表
// https://work.weixin.qq.com/api/doc/90000/90135/92120
func (comp *GroupChat) GroupChatList(params *request.RequestAddGroupChat) (*response.ResponseGroupChatList, error) {

	result := &response.ResponseGroupChatList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/list", params, nil, nil, result)

	return result, err
}

// 获取客户群详情
// https://work.weixin.qq.com/api/doc/90000/90135/92122
func (comp *GroupChat) GetGroupChat(chatID string, needName bool) (*response.ResponseGroupChatGet, error) {

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
func (comp *GroupChat) OpenGIDToChatID(openGID string) (*response.ResponseGroupChatOpenGIDToChatID, error) {

	result := &response.ResponseGroupChatOpenGIDToChatID{}

	options := &object.HashMap{
		"opengid": openGID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/opengid_to_chatid?", options, nil, nil, result)

	return result, err
}
