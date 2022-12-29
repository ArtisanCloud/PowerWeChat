package customer

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/customer/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/customer/response"
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

// 读取消息
// https://developer.work.weixin.qq.com/document/path/94670
func (comp *Client) BatchGet(ctx context.Context, externalUserIDList []string) (*response.ResponseCustomerBatchGet, error) {

	result := &response.ResponseCustomerBatchGet{}

	options := &object.HashMap{
		"external_userid_list": externalUserIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/customer/batchget", options, nil, nil, result)

	return result, err
}

// 获取配置的专员与客户群
// https://developer.work.weixin.qq.com/document/path/94674
func (comp *Client) GetUpgradeServiceConfig(ctx context.Context) (*response.ResponseCustomerGetUpgradeServiceConfig, error) {

	result := &response.ResponseCustomerGetUpgradeServiceConfig{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/kf/customer/get_upgrade_service_config", nil, nil, result)

	return result, err
}

// 为客户升级为专员或客户群服务
// https://developer.work.weixin.qq.com/document/path/94674
func (comp *Client) UpgradeService(ctx context.Context, options *request.RequestUpgradeService) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/customer/upgrade_service", options, nil, nil, result)

	return result, err
}

// 为客户取消推荐
// https://developer.work.weixin.qq.com/document/path/94674
func (comp *Client) CancelUpgradeService(ctx context.Context, openKFID, externalUserID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}
	options := &power.StringMap{
		"open_kfid":       openKFID,
		"external_userid": externalUserID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/customer/cancel_upgrade_service", options, nil, nil, result)

	return result, err
}
