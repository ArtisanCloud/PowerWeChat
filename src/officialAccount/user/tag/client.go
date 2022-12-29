package tag

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user/tag/response"
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

// 创建标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Create(ctx *context.Context, name string) (*response.ResponseTagGet, error) {
	result := &response.ResponseTagGet{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"name": name,
		},
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/create", params, nil, nil, result)

	return result, err
}

// 获取公众号已创建的标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) List(ctx *context.Context) (*response.ResponseTagGetList, error) {

	result := &response.ResponseTagGetList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/tags/get", nil, nil, result)

	return result, err

}

// 用户标签管理
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Update(ctx *context.Context, tagID string, name string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"id":   tagID,
			"name": name,
		},
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/update", params, nil, nil, result)

	return result, err
}

// 删除标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Delete(ctx *context.Context, tagID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"id": tagID,
		},
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/delete", params, nil, nil, result)

	return result, err
}

// 获取用户身上的标签列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UserTags(ctx *context.Context, openID string) (*response.ResponseUserTags, error) {
	result := &response.ResponseUserTags{}

	params := &object.HashMap{
		"openid": openID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/getidlist", params, nil, nil, result)

	return result, err
}

// 获取标签下粉丝列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UsersOfTag(ctx *context.Context, tagID string, nextOpenID string) (*response.ResponseUserOfTag, error) {
	result := &response.ResponseUserOfTag{}

	params := &object.HashMap{
		"tagid":       tagID,
		"next_openid": nextOpenID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/user/tag/get", params, nil, nil, result)

	return result, err
}

// 批量为用户打标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) TagUsers(ctx *context.Context, openIDs []string, tagID string) (*response.ResponseTagUsers, error) {
	result := &response.ResponseTagUsers{}

	params := &object.HashMap{
		"openid_list": openIDs,
		"tagid":       tagID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/members/batchtagging", params, nil, nil, result)

	return result, err
}

// 批量为用户取消标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UntagUsers(ctx *context.Context, openIDs []string, tagID string) (*response.ResponseUntagUsers, error) {
	result := &response.ResponseUntagUsers{}

	params := &object.HashMap{
		"openid_list": openIDs,
		"tagid":       tagID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/tags/members/batchuntagging", params, nil, nil, result)

	return result, err
}
