package user

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	response3 "github.com/ArtisanCloud/power-wechat/src/work/user/request"
	"github.com/ArtisanCloud/power-wechat/src/work/user/response"
	"github.com/golang-module/carbon"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90194

// 创建成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90195
func (comp *Client) Create(data *response3.RequestUserDetail) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/user/create", data, nil, result)

	return result
}

// 更新成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90197
func (comp *Client) Update(data *response3.RequestUserDetail) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/user/update", data, nil, result)

	return result
}

// 删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func (comp *Client) Delete(userID string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpGet("cgi-bin/user/delete", &object.StringMap{
		"userid": userID,
	}, result)

	return result
}

// 批量删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func (comp *Client) BatchDelete(userIDs []string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/user/batchdelete", &object.HashMap{
		"useridlist": userIDs,
	}, nil, result)

	return result
}

// 获取成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func (comp *Client) Get(userID string) *response.ResponseGetUserDetail {

	result := &response.ResponseGetUserDetail{}

	comp.HttpGet("cgi-bin/user/get", &object.StringMap{
		"userid": userID,
	}, result)

	return result
}

// 获取部门的成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90200
func (comp *Client) GetDepartmentUsers(departmentID int, fetchChild int) *response.ResponseGetSimpleUserList {

	result := &response.ResponseGetSimpleUserList{}

	comp.HttpGet("cgi-bin/user/simplelist", &object.StringMap{
		"department_id": fmt.Sprintf("%d", departmentID),
		"fetch_child":   fmt.Sprintf("%d", fetchChild),
	}, result)

	return result
}

// 获取部门成员详情
// https://open.work.weixin.qq.com/api/doc/90000/90135/90201
func (comp *Client) GetDetailedDepartmentUsers(departmentID int, fetchChild int) *response.ResponseGetUserList {

	result := &response.ResponseGetUserList{}

	comp.HttpGet("cgi-bin/user/list", &object.StringMap{
		"department_id": fmt.Sprintf("%d", departmentID),
		"fetch_child":   fmt.Sprintf("%d", fetchChild),
	}, result)

	return result
}

// userid与openid互换
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func (comp *Client) UserIdToOpenid(userID string) *response.ResponseUserIDToOpenID {

	result := &response.ResponseUserIDToOpenID{}

	comp.HttpPostJson("cgi-bin/user/convert_to_openid", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result
}

// userid与openid互换
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func (comp *Client) OpenIDToUserID(openID string) *response.ResponseOpenIDToUserID {

	result := &response.ResponseOpenIDToUserID{}

	comp.HttpPostJson("cgi-bin/user/convert_to_userid", &object.StringMap{
		"openid": openID,
	}, nil, result)

	return result
}

// userid与openid互换
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func (comp *Client) MobileToUserID(mobile string) *response.ResponseMobileToUserID {

	result := &response.ResponseMobileToUserID{}

	comp.HttpPostJson("cgi-bin/user/getuserid", &object.StringMap{
		"mobile": mobile,
	}, nil, result)

	return result
}

// 二次验证
// https://open.work.weixin.qq.com/api/doc/90000/90135/90203
func (comp *Client) Accept(userID string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpGet("cgi-bin/user/authsucc", &object.StringMap{
		"userid": userID,
	}, result)

	return result
}

// 邀请成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90975
func (comp *Client) Invite(params *object.HashMap) *response.ResponseMobileToUserID {

	result := &response.ResponseMobileToUserID{}

	comp.HttpPostJson("cgi-bin/batch/invite", params, nil, result)

	return result
}

// 获取加入企业二维码
// https://open.work.weixin.qq.com/api/doc/90000/90135/91714
func (comp *Client) GetInvitationQrCode(sizeType int) (*response.ResponseJoinCode, error) {

	if sizeType < 1 || sizeType > 4 {
		return nil, errors.New("The sizeType must be 1, 2, 3, 4.")
	}

	result := &response.ResponseJoinCode{}

	comp.HttpGet("cgi-bin/corp/get_join_qrcode", &object.StringMap{
		"size_type": fmt.Sprintf("%d", sizeType),
	}, result)

	return result, nil
}


// 获取企业活跃成员数
// https://open.work.weixin.qq.com/api/doc/90000/90135/92714
func (comp *Client) GetActiveCount(date carbon.Carbon) (*response.ResponseUserActiveCount) {

	result := &response.ResponseUserActiveCount{}

	comp.HttpPostJson("cgi-bin/user/get_active_stat", &object.HashMap{
		"date": date,
	},nil, result)

	return result
}