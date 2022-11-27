package msgaudit

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/msgAudit/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取会话内容存档开启成员列表
// https://developer.work.weixin.qq.com/document/path/91614
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
// https://developer.work.weixin.qq.com/document/path/91782
func (comp *Client) CheckSingleAgree(info []*power.StringMap) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"info": info,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_single_agree", params, nil, nil, result)

	return result, err
}

// 获取会话同意情况
// https://developer.work.weixin.qq.com/document/path/91782
func (comp *Client) CheckRoomAgree(roomID string) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_room_agree?", params, nil, nil, result)

	return result, err
}

// 获取会话内容存档内部群信息
// https://developer.work.weixin.qq.com/document/path/92951
func (comp *Client) GroupchatGet(roomID string) (*response.ResponseMsgAuditGetRoom, error) {

	result := &response.ResponseMsgAuditGetRoom{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/groupchat/get?", params, nil, nil, result)

	return result, err
}

// 获取机器人信息
// https://developer.work.weixin.qq.com/document/path/91774
func (comp *Client) GetRobotInfo(robotID string) (*response.ResponseMsgAuditGetRobotInfo, error) {

	result := &response.ResponseMsgAuditGetRobotInfo{}

	params := &object.StringMap{
		"robot_id": robotID,
	}

	_, err := comp.HttpGet("cgi-bin/msgaudit/get_robot_info?", params, nil, result)

	return result, err
}
