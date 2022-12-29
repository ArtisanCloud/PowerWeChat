package menu

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/menu/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/menu/response"
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

// 获取自定义菜单配置
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Getting_Custom_Menu_Configurations.html
func (comp *Client) Get(ctx *context.Context) (*response.ResponseMenuGet, error) {

	result := &response.ResponseMenuGet{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/menu/get", nil, nil, result)

	return result, err
}

// 查询接口
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html
func (comp *Client) CurrentSelfMenu(ctx *context.Context) (*response.ResponseCurrentSelfMenu, error) {

	result := &response.ResponseCurrentSelfMenu{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/get_current_selfmenu_info", nil, nil, result)

	return result, err
}

// 创建接口
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
func (comp *Client) Create(ctx *context.Context, buttons []*request.Button) (*response.ResponseMenuCreate, error) {

	result := &response.ResponseMenuCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/menu/create", &object.HashMap{
		"button": buttons,
	}, nil, nil, result)

	return result, err
}

// 创建个性化菜单
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html
func (comp *Client) CreateConditional(ctx *context.Context, buttons []*request.Button, rules *request.RequestMatchRule) (*response.ResponseMenuCreateConditional, error) {

	result := &response.ResponseMenuCreateConditional{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/menu/addconditional", &object.HashMap{
		"button":    buttons,
		"matchrule": rules,
	}, nil, nil, result)

	return result, err

}

// 删除接口
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html
func (comp *Client) Delete(ctx *context.Context) (*response.ResponseMenuDelete, error) {

	result := &response.ResponseMenuDelete{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/menu/delete", nil, nil, result)

	return result, err
}

// 删除个性化菜单
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html
func (comp *Client) DeleteConditional(ctx *context.Context, menuID int) (*response.ResponseMenuDelete, error) {

	result := &response.ResponseMenuDelete{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/menu/delconditional", &object.HashMap{
		"menuid": menuID,
	}, nil, nil, result)

	return result, err
}

// 测试个性化菜单匹配结果
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html#1
func (comp *Client) TryMatch(ctx *context.Context, userID string) (*response.ResponseMenuTryMatch, error) {

	result := &response.ResponseMenuTryMatch{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/menu/trymatch", &object.HashMap{
		"user_id": userID,
	}, nil, nil, result)

	return result, err
}
