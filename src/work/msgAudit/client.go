package msgaudit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/msgAudit/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取会话内容存档开启成员列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/91614
func (comp *Client) GetPermitUsersList(msgType string) (*response.ResponseMsgAuditGetPermitUsers, error) {

	result := &response.ResponseMsgAuditGetPermitUsers{}

	params := &object.StringMap{}
	if msgType != "" {
		(*params)["type"] = msgType
	}
	_, err := comp.HttpPostJson("cgi-bin/msgaudit/get_permit_user_list", params, nil, nil, result)

	return result, err
}

// 获取会话同意情况
// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func (comp *Client) CheckSingleAgree(info []*power.StringMap) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"info": info,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_single_agree", params, nil, nil, result)

	return result, err
}

// 获取会话同意情况
// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func (comp *Client) CheckRoomAgree(roomID string) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_room_agree?", params, nil, nil, result)

	return result, err
}

// 获取会话内容存档内部群信息
// https://open.work.weixin.qq.com/api/doc/90000/90135/92951
func (comp *Client) GroupchatGet(roomID string) (*response.ResponseMsgAuditGetRoom, error) {

	result := &response.ResponseMsgAuditGetRoom{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/groupchat/get?", params, nil, nil, result)

	return result, err
}
