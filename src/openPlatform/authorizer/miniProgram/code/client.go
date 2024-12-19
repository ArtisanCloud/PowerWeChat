package code

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	response4 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/code/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/code/response"
	"net/http"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 上传小程序代码并生成体验版
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html
func (comp *Client) Commit(ctx context.Context, templateID string, extJson string, version string, description string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"template_id":  templateID,
		"ext_json":     extJson,
		"user_version": version,
		"user_desc":    description,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/commit", params, nil, nil, result)

	return result, err

}

// 获取体验版二维码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (comp *Client) GetQrCode(ctx context.Context, path string) (*http.Response, error) {

	var header = &response4.ResponseHeaderMedia{}

	rs, err := comp.BaseClient.RequestRaw(ctx, "wxa/get_qrcode", http.MethodGet, &object.HashMap{
		"query": &object.StringMap{
			"path": path,
		}}, &header, nil)

	return rs, err

}

// 获取审核时可填写的类目信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html#请求地址
func (comp *Client) GetCategory(ctx context.Context) (*response.ResponseGetCategory, error) {

	result := &response.ResponseGetCategory{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_category", nil, nil, result)

	return result, err

}

// 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html#请求地址
func (comp *Client) GetPage(ctx context.Context) (*response.ResponseGetPage, error) {

	result := &response.ResponseGetPage{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_page", nil, nil, result)

	return result, err

}

// 提交审核
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (comp *Client) SubmitAudit(ctx context.Context, params *request.RequestSubmitAudit) (*response.ResponseSubmitAudit, error) {

	result := &response.ResponseSubmitAudit{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/submit_audit", params, nil, nil, result)

	return result, err

}

// 查询指定发布审核单的审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html#请求地址
func (comp *Client) GetAuditStatus(ctx context.Context, auditID int) (*response.ResponseGetAuditStatus, error) {

	result := &response.ResponseGetAuditStatus{}

	params := &object.HashMap{
		"auditid": auditID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/get_auditstatus", params, nil, nil, result)

	return result, err

}

// 查询最新一次提交的审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_latest_auditstatus.html#请求地址
func (comp *Client) GetLatestAuditStatus(ctx context.Context) (*response.ResponseGetLatestAuditStatus, error) {

	result := &response.ResponseGetLatestAuditStatus{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_latest_auditstatus", nil, nil, result)

	return result, err

}

// 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html#请求地址
func (comp *Client) Release(ctx context.Context) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/release", nil, nil, nil, result)

	return result, err

}

// 小程序审核撤回
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/undocodeaudit.html#请求地址
func (comp *Client) WithdrawAudit(ctx context.Context) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/undocodeaudit", nil, nil, result)

	return result, err

}

// 获取可回退的小程序版本
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_history_version.html#请求地址
func (comp *Client) RollbackRelease(ctx context.Context) (*response.ResponseRollbackRelease, error) {

	result := &response.ResponseRollbackRelease{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/revertcoderelease", nil, nil, result)

	return result, err

}

// 修改小程序服务状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/change_visitstatus.html#请求地址
func (comp *Client) ChangeVisitStatus(ctx context.Context, action string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"action": action,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/change_visitstatus", params, nil, nil, result)

	return result, err

}

// 分阶段发布
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/grayrelease.html#请求地址
func (comp *Client) GrayRelease(ctx context.Context, grayPercentage int) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"gray_percentage": grayPercentage,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/grayrelease", params, nil, nil, result)

	return result, err

}

// 取消分阶段发布
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertgrayrelease.html#请求地址
func (comp *Client) RevertGrayRelease(ctx context.Context) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/revertgrayrelease", nil, nil, result)

	return result, err

}

// 查询当前分阶段发布详情
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html#请求地址
func (comp *Client) GetGrayRelease(ctx context.Context) (*response.ResponseGetGrayRelease, error) {

	result := &response.ResponseGetGrayRelease{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/getgrayreleaseplan", nil, nil, result)

	return result, err

}

// 查询当前设置的最低基础库版本及各版本用户占比
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getweappsupportversion.html#请求地址
func (comp *Client) GetSupportVersion(ctx context.Context) (*response.ResponseGetSupportVersion, error) {

	result := &response.ResponseGetSupportVersion{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxopen/getweappsupportversion", nil, nil, nil, result)

	return result, err

}

// 设置最低基础库版本
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/setweappsupportversion.html#请求地址
func (comp *Client) SetSupportVersion(ctx context.Context, version string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxopen/setweappsupportversion", &object.HashMap{
		"version": version,
	}, nil, nil, result)

	return result, err

}

// 查询服务商的当月提审限额（quota）和加急次数
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/query_quota.html#请求地址
func (comp *Client) QueryQuota(ctx context.Context) (*response.ResponseQueryQuota, error) {

	result := &response.ResponseQueryQuota{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/queryquota", nil, nil, result)

	return result, err

}

// 加急审核申请
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html#请求地址
func (comp *Client) SpeedupAudit(ctx context.Context, auditID int) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/speedupaudit", &object.HashMap{
		"auditid": auditID,
	}, nil, nil, result)

	return result, err

}
