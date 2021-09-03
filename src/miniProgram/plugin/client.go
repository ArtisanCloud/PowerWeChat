package plugin

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/plugin/response"
)

type Client struct {
	*kernel.BaseClient
}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html
func (comp *Client) CreateActivityID(action string, pluginAppID string, reason string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action":       "apply",
		"plugin_appid": pluginAppID,
		"reason":       reason,
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 查询已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html
func (comp *Client) GetPluginList() (*response.ResponsePluginGetPluginList, error) {

	result := &response.ResponsePluginGetPluginList{}

	data := &object.HashMap{
		"action": "list",
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 删除已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.unbindPlugin.html
func (comp *Client) UnbindPlugin(pluginAppID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action":       "unbind",
		"plugin_appid": pluginAppID,
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/plugin", data, nil, nil, result)

	return result, err
}

// 获取当前所有插件使用方（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html
func (comp *Client) GetPluginDevApplyList(action string, page int, num int) (*response.ResponsePluginGetDevApplyList, error) {

	result := &response.ResponsePluginGetDevApplyList{}

	data := &object.HashMap{
		"action": action,
		"page":   page,
		"num":    num,
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/devplugin", data, nil, nil, result)

	return result, err
}

// 修改插件使用申请的状态（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.setDevPluginApplyStatus.html
func (comp *Client) SetDevPluginApplyStatus(action string, appID string, reason string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"action": action,
		"appid":  appID,
		"reason": reason,
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/devplugin", data, nil, nil, result)

	return result, err
}
