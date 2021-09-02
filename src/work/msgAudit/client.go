package msgaudit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/msgaudit/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91614
func (comp *Client) GetPermitUsers(msgType string) (*response.ResponseMsgAuditGetPermitUsers, error) {

	result := &response.ResponseMsgAuditGetPermitUsers{}

	params := &object.StringMap{}
	if msgType != "" {
		(*params)["type"] = msgType
	}
	_, err := comp.HttpPostJson("cgi-bin/msgaudit/get_permit_user_list", params, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func (comp *Client) GetSingleAgreeStatus(info *power.HashMap) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"info": info.ToHashMap(),
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_single_agree", params, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func (comp *Client) GetRoomAgreeStatus(roomID string) (*response.ResponseMsgAuditGetAgreeInfo, error) {

	result := &response.ResponseMsgAuditGetAgreeInfo{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/check_room_agree?", params, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/92951
func (comp *Client) GetRoom(roomID string) (*response.ResponseMsgAuditGetRoom, error) {

	result := &response.ResponseMsgAuditGetRoom{}

	params := &object.HashMap{
		"roomid": roomID,
	}

	_, err := comp.HttpPostJson("cgi-bin/msgaudit/groupchat/get?", params, nil, nil, result)

	return result, err
}
