package user

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user/response"
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

// 获取用户基本信息
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (comp *Client) Get(openID string, lang string) (*response.ResponseGetUserInfo, error) {

	result := &response.ResponseGetUserInfo{}

	params := &object.StringMap{
		"openid": openID,
		"lang":   lang,
	}

	_, err := comp.HttpGet("cgi-bin/user/info", params, nil, result)

	return result, err

}

// 批量获取用户基本信息
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (comp *Client) BatchGet(data *request.RequestBatchGetUserInfo) (*response.ResponseBatchGetUserInfo, error) {

	result := &response.ResponseBatchGetUserInfo{}

	_, err := comp.HttpPostJson("cgi-bin/user/info/batchget", data, nil, nil, result)

	return result, err
}

// 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (comp *Client) List(nextOpenID string) (*response.ResponseGetUserList, error) {

	result := &response.ResponseGetUserList{}

	params := &object.StringMap{
		"next_openid": nextOpenID,
	}

	_, err := comp.HttpGet("cgi-bin/user/get", params, nil, result)

	return result, err

}

// 设置用户备注名
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html
func (comp *Client) Remark(openID string, remark string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.StringMap{
		"openid": openID,
		"remark": remark,
	}

	_, err := comp.HttpPostJson("cgi-bin/user/info/updateremark", params, nil, nil, result)

	return result, err
}

// 黑名单管理
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (comp *Client) Blacklist(beginOpenID string) (*response.ResponseBlacklist, error) {

	result := &response.ResponseBlacklist{}

	params := &object.StringMap{
		"begin_openid": beginOpenID,
	}

	_, err := comp.HttpPostJson("cgi-bin/tags/members/getblacklist", params, nil, nil, result)

	return result, err
}

// 拉黑用户
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (comp *Client) Block(openIDList []string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"openid_list": openIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/tags/members/batchblacklist", params, nil, nil, result)

	return result, err
}

// 取消拉黑用户
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (comp *Client) Unblock(openIDList []string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"openid_list": openIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/tags/members/batchunblacklist", params, nil, nil, result)

	return result, err
}

// 帐号迁移 转换 openid
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (comp *Client) ChangeOpenID(fromAppID string, openIDList []string) (*response.ResponseChangeOpenID, error) {

	result := &response.ResponseChangeOpenID{}

	params := &object.HashMap{
		"from_appid":  fromAppID,
		"openid_list": openIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/changeopenid", params, nil, nil, result)

	return result, err
}
