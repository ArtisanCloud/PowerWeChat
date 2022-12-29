package plugin

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/plugin/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html
func (comp *Client) ApplyPlugin(ctx context.Context, pluginAppID string, reason string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action":       "apply",
		"plugin_appid": pluginAppID,
		"reason":       reason,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 查询已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html
func (comp *Client) GetPluginList(ctx context.Context) (*response.ResponsePluginGetPluginList, error) {

	result := &response.ResponsePluginGetPluginList{}

	data := &object.HashMap{
		"action": "list",
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 删除已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.unbindPlugin.html
func (comp *Client) UnbindPlugin(ctx context.Context, pluginAppID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action":       "unbind",
		"plugin_appid": pluginAppID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 获取当前所有插件使用方（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html
func (comp *Client) GetPluginDevApplyList(ctx context.Context, action string, page int, num int) (*response.ResponsePluginGetDevApplyList, error) {

	result := &response.ResponsePluginGetDevApplyList{}

	data := &object.HashMap{
		"action": action,
		"page":   page,
		"num":    num,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxa/devplugin", data, nil, nil, result)

	return result, err
}

// 修改插件使用申请的状态（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.setDevPluginApplyStatus.html
func (comp *Client) SetDevPluginApplyStatus(ctx context.Context, action string, appID string, reason string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action": action,
		"appid":  appID,
		"reason": reason,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wxa/devplugin", data, nil, nil, result)

	return result, err
}
