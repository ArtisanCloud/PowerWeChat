package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

// 创建成员
// https://developer.work.weixin.qq.com/document/path/90195
func (comp *Client) Create(ctx context.Context, data *response3.RequestUserDetail) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/create", data, nil, nil, result)

	return result, err
}

// 更新成员
// https://developer.work.weixin.qq.com/document/path/90197
func (comp *Client) Update(ctx context.Context, data *response3.RequestUserDetail) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/update", data, nil, nil, result)

	return result, err
}

// 删除成员
// https://developer.work.weixin.qq.com/document/path/90198
func (comp *Client) Delete(ctx context.Context, userID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/user/delete", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result, err
}

// 批量删除成员
// https://developer.work.weixin.qq.com/document/path/90335
func (comp *Client) BatchDelete(ctx context.Context, userIDs []string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/batchdelete", &object.HashMap{
		"useridlist": userIDs,
	}, nil, nil, result)

	return result, err
}

// 获取成员
// https://developer.work.weixin.qq.com/document/path/90196
func (comp *Client) Get(ctx context.Context, userID string) (*response.ResponseGetUserDetail, error) {

	result := &response.ResponseGetUserDetail{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/user/get", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result, err
}

// 获取部门的成员
// https://developer.work.weixin.qq.com/document/path/90200
func (comp *Client) GetDepartmentUsers(ctx context.Context, departmentID int, fetchChild int) (*response.ResponseGetSimpleUserList, error) {

	result := &response.ResponseGetSimpleUserList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/user/simplelist", &object.StringMap{
		"department_id": fmt.Sprintf("%d", departmentID),
		"fetch_child":   fmt.Sprintf("%d", fetchChild),
	}, nil, result)

	return result, err
}

// 获取部门成员详情
// https://developer.work.weixin.qq.com/document/path/90201
func (comp *Client) GetDetailedDepartmentUsers(ctx context.Context, departmentID int, fetchChild int) (*response.ResponseGetUserList, error) {

	result := &response.ResponseGetUserList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/user/list", &object.StringMap{
		"department_id": fmt.Sprintf("%d", departmentID),
		"fetch_child":   fmt.Sprintf("%d", fetchChild),
	}, nil, result)

	return result, err
}

// userid与openid互换
// https://developer.work.weixin.qq.com/document/path/90202
func (comp *Client) UserIdToOpenID(ctx context.Context, userID string) (*response.ResponseUserIDToOpenID, error) {

	result := &response.ResponseUserIDToOpenID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/convert_to_openid", &object.StringMap{
		"userid": userID,
	}, nil, nil, result)

	return result, err
}

// 获取成员ID列表
// https://developer.work.weixin.qq.com/document/path/96067
func (comp *Client) ListID(ctx context.Context, cursor string, limit int) (*response.ResponseListID, error) {

	result := &response.ResponseListID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/list_id", &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}, nil, nil, result)

	return result, err
}

// openid转userid
// https://developer.work.weixin.qq.com/document/path/90202
func (comp *Client) OpenIDToUserID(ctx context.Context, openID string) (*response.ResponseOpenIDToUserID, error) {

	result := &response.ResponseOpenIDToUserID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/convert_to_userid", &object.StringMap{
		"openid": openID,
	}, nil, nil, result)

	return result, err
}

// 手机号获取userid
// https://developer.work.weixin.qq.com/document/path/95402
func (comp *Client) MobileToUserID(ctx context.Context, mobile string) (*response.ResponseMobileToUserID, error) {

	result := &response.ResponseMobileToUserID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/getuserid", &object.StringMap{
		"mobile": mobile,
	}, nil, nil, result)

	return result, err
}

// 邮箱获取userid
// https://developer.work.weixin.qq.com/document/path/95895
func (comp *Client) EmailToUserID(ctx context.Context, email string, emailType int) (*response.ResponseConvertToUserID, error) {

	result := &response.ResponseConvertToUserID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/get_userid_by_email", &object.HashMap{
		"email":      email,
		"email_type": emailType,
	}, nil, nil, result)

	return result, err
}

// 二次验证
// https://developer.work.weixin.qq.com/document/path/90203
func (comp *Client) Accept(ctx context.Context, userID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/user/authsucc", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result, err
}

// 邀请成员
// https://developer.work.weixin.qq.com/document/path/90975
func (comp *Client) Invite(ctx context.Context, params *power.HashMap) (*response.ResponseMobileToUserID, error) {

	result := &response.ResponseMobileToUserID{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/invite", params, nil, nil, result)

	return result, err
}

// 获取加入企业二维码
// https://developer.work.weixin.qq.com/document/path/91714
func (comp *Client) GetJoinQrCode(ctx context.Context, sizeType int) (*response.ResponseJoinCode, error) {

	if sizeType < 1 || sizeType > 4 {
		return nil, errors.New("The sizeType must be 1, 2, 3, 4.")
	}

	result := &response.ResponseJoinCode{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/corp/get_join_qrcode", &object.StringMap{
		"size_type": fmt.Sprintf("%d", sizeType),
	}, nil, result)

	return result, err
}

// 获取企业活跃成员数
// https://developer.work.weixin.qq.com/document/path/92714
func (comp *Client) GetActiveStat(ctx context.Context, date string) (*response.ResponseUserActiveCount, error) {

	result := &response.ResponseUserActiveCount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/get_active_stat", &object.HashMap{
		"date": date,
	}, nil, nil, result)

	return result, err
}
