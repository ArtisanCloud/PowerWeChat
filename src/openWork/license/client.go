package license

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	kernelResponse "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/provider"
)

type Client struct {
	BaseClient *kernel.BaseClient
	corpID     string
}

func NewClient(app kernel.ApplicationInterface, corpID string) (*Client, error) {
	token, err := provider.NewAccessToken(&app, corpID)
	if err != nil {
		return nil, err
	}
	baseClient, err := kernel.NewBaseClient(&app, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
		corpID,
	}, nil
}

// 下单购买账号
// https://developer.work.weixin.qq.com/document/path/95644
func (comp *Client) CreateNewOrder(ctx context.Context, req *request.RequestCreateNewOrder) (string, error) {

	var result struct {
		kernelResponse.ResponseWork
		OrderID string `json:"order_id,omitempty"`
	}

	req.CorpID = comp.corpID

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/create_new_order", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return "", result
	}

	return result.OrderID, err
}

// 创建续期任务
// https://developer.work.weixin.qq.com/document/path/95646
func (comp *Client) CreateRenewOrderJob(ctx context.Context, req *request.RequestCreateRenewOrderJob) (*response.ResponseCreateRenewOrderJob, error) {
	var result response.ResponseCreateRenewOrderJob

	req.CorpID = comp.corpID

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/create_renew_order_job", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}
	return &result, err
}

// 提交续期订单
// https://developer.work.weixin.qq.com/document/path/95646#%E6%8F%90%E4%BA%A4%E7%BB%AD%E6%9C%9F%E8%AE%A2%E5%8D%95
func (comp *Client) SubmitOrderJob(ctx context.Context, req *request.RequestSubmitOrderJob) (string, error) {

	var result struct {
		kernelResponse.ResponseWork
		OrderID string `json:"order_id,omitempty"`
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/submit_order_job", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return "", result
	}

	return result.OrderID, err
}

// 获取订单列表
// https://developer.work.weixin.qq.com/document/path/95647
func (comp *Client) ListOrder(ctx context.Context, req *request.RequestListOrder) (*response.ResponseListOrder, error) {
	var result response.ResponseListOrder
	req.CorpID = comp.corpID

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/list_order", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return &result, err
}

// 获取订单详情
// https://developer.work.weixin.qq.com/document/path/95648
func (comp *Client) GetOrder(ctx context.Context, orderID string) (*model.Order, error) {
	var result response.ResponseGetOrder

	req := object.HashMap{
		"order_id": orderID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/get_order", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return result.Order, err
}

// 获取订单中的账号列表
// https://developer.work.weixin.qq.com/document/path/95649
func (comp *Client) ListOrderAccount(ctx context.Context, orderID string, limit int, cursor string) (*response.ResponseListOrderAccount, error) {

	var result response.ResponseListOrderAccount

	var req = object.HashMap{
		"order_id": orderID,
	}
	if limit > 0 {
		req["limit"] = limit
	}
	if cursor != "" {
		req["cursor"] = cursor
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/list_order_account", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return &result, err
}

// 取消订单
// https://developer.work.weixin.qq.com/document/path/96106
func (comp *Client) CancelOrder(ctx context.Context, orderID string) error {
	var result kernelResponse.ResponseWork

	req := object.HashMap{
		"corpid":   comp.corpID,
		"order_id": orderID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/cancel_order", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return result
	}

	return err
}

// 激活账号
// https://developer.work.weixin.qq.com/document/path/95553#%E6%BF%80%E6%B4%BB%E5%B8%90%E5%8F%B7
func (comp *Client) ActiveAccount(ctx context.Context, userID string, activeCode string) error {
	var result kernelResponse.ResponseWork

	req := model.ActiveInfo{
		CorpID:     comp.corpID,
		UserID:     userID,
		ActiveCode: activeCode,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/active_account", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return result
	}

	return err
}

// 批量激活账号
// https://qyapi.weixin.qq.com/cgi-bin/license/batch_active_account?provider_access_token=ACCESS_TOKEN
func (comp *Client) BatchActiveAccount(ctx context.Context, list []model.ActiveInfo) error {
	var result kernelResponse.ResponseWork
	req := object.HashMap{
		"corpid":      comp.corpID,
		"active_list": list,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/batch_active_account", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return result
	}

	return err
}

// 获取激活码详情
// https://developer.work.weixin.qq.com/document/path/95552#%E8%8E%B7%E5%8F%96%E6%BF%80%E6%B4%BB%E7%A0%81%E8%AF%A6%E6%83%85
func (comp *Client) GetActiveInfoByCode(ctx context.Context, activeCode string) (*model.ActiveInfo, error) {
	var result response.ResponseGetActiveInfoByCode
	req := object.HashMap{
		"corpid":      comp.corpID,
		"active_code": activeCode,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/get_active_info_by_code", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return result.ActiveInfo, err
}

// 批量获取激活码详情
// https://developer.work.weixin.qq.com/document/path/95552#%E6%89%B9%E9%87%8F%E8%8E%B7%E5%8F%96%E6%BF%80%E6%B4%BB%E7%A0%81%E8%AF%A6%E6%83%85
func (comp *Client) BatchGetActiveInfoByCode(ctx context.Context, activeCodeList []string) ([]model.ActiveInfo, error) {
	var result response.ResponseBatchGetActiveInfoByCode
	req := object.HashMap{
		"corpid":           comp.corpID,
		"active_code_list": activeCodeList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/batch_get_active_info_by_code", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return result.ActiveInfoList, err
}

// 获取企业的账号列表
// https://developer.work.weixin.qq.com/document/path/95544
func (comp *Client) ListActivatedAccount(ctx context.Context, limit int, cursor string) (*response.ResponseListActivatedAccount, error) {
	var result response.ResponseListActivatedAccount
	req := object.HashMap{
		"corpid": comp.corpID,
	}

	if limit > 0 {
		req["limit"] = limit
	}
	if cursor != "" {
		req["cursor"] = cursor
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/list_activated_account", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return &result, err
}

// 获取成员的激活详情
// https://developer.work.weixin.qq.com/document/path/95555
func (comp *Client) GetActiveInfoByUser(ctx context.Context, userID string) (*response.ResponseGetActiveInfoByUser, error) {
	var result response.ResponseGetActiveInfoByUser
	req := object.HashMap{
		"corpid": comp.corpID,
		"userid": userID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/get_active_info_by_user", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return &result, err
}

// 账号继承
// https://developer.work.weixin.qq.com/document/path/95673
func (comp *Client) BatchTransferLicense(ctx context.Context, transferList []model.TransferInfo) ([]model.TransferInfo, error) {
	var result struct {
		kernelResponse.ResponseWork
		TransferResult []model.TransferInfo `json:"transfer_result,omitempty"`
	}
	req := object.HashMap{
		"corpid":        comp.corpID,
		"transfer_list": transferList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/batch_transfer_license", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return result.TransferResult, err
}

// 获取应用的接口许可状态
// https://developer.work.weixin.qq.com/document/path/95844
func (comp *Client) GetAppLicenseInfo(ctx context.Context) (*model.LicenseInfo, error) {
	var result response.ResponseGetAppLicenseInfo
	config := (*comp.BaseClient.App).GetContainer().GetConfig()
	req := object.HashMap{
		"corpid":   comp.corpID,
		"suite_id": (*config)["appid"].(string),
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/get_app_license_info", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return nil, result
	}

	return &result.LicenseInfo, err
}

// 设置企业的许可自动激活状态
// https://developer.work.weixin.qq.com/document/path/95873
func (comp *Client) SetAutoActiveStatus(ctx context.Context, autoActiveStatus int) error {
	var result kernelResponse.ResponseWork
	req := object.HashMap{
		"corpid":             comp.corpID,
		"auto_active_status": autoActiveStatus,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/set_auto_active_status", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return result
	}

	return err
}

// 查询企业的许可自动激活状态
// https://developer.work.weixin.qq.com/document/path/95874
func (comp *Client) GetAutoActiveStatus(ctx context.Context) (int, error) {
	var result struct {
		kernelResponse.ResponseWork
		AutoActiveStatus int `json:"auto_active_status,omitempty"`
	}
	req := object.HashMap{
		"corpid": comp.corpID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/license/get_auto_active_status", &req, nil, nil, &result)
	if err == nil && result.IsError() {
		return 0, result
	}

	return result.AutoActiveStatus, err
}
