package tag

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/user/tag/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 创建标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Create(name string) (*response.ResponseTagGet, error) {
	result := &response.ResponseTagGet{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"name": name,
		},
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/create", params, nil, nil, result)

	return result, err
}

// 获取公众号已创建的标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) List() (*response.ResponseTagGetList, error) {

	result := &response.ResponseTagGetList{}

	_, err := comp.HttpGet("cgi-bin/tags/get", nil, nil, result)

	return result, err

}

// 用户标签管理
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Update(tagID string, name string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"id":   tagID,
			"name": name,
		},
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/update", params, nil, nil, result)

	return result, err
}

// 删除标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) Delete(tagID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag": &object.HashMap{
			"id": tagID,
		},
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/delete", params, nil, nil, result)

	return result, err
}

// 获取用户身上的标签列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UserTags(openID string) (*response.ResponseUserTags, error) {
	result := &response.ResponseUserTags{}

	params := &object.HashMap{
		"openid": openID,
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/getidlist", params, nil, nil, result)

	return result, err
}

// 获取标签下粉丝列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UsersOfTag(tagID string, nextOpenID string) (*response.ResponseUserOfTag, error) {
	result := &response.ResponseUserOfTag{}

	params := &object.HashMap{
		"tagid":       tagID,
		"next_openid": nextOpenID,
	}
	_, err := comp.HttpPostJson("cgi-bin/user/tag/get", params, nil, nil, result)

	return result, err
}

// 批量为用户打标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) TagUsers(openIDs []string, tagID string) (*response.ResponseTagUsers, error) {
	result := &response.ResponseTagUsers{}

	params := &object.HashMap{
		"openid_list": openIDs,
		"tagid":       tagID,
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/members/batchtagging", params, nil, nil, result)

	return result, err
}

// 批量为用户取消标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (comp *Client) UntagUsers(openIDs []string, tagID string) (*response.ResponseUntagUsers, error) {
	result := &response.ResponseUntagUsers{}

	params := &object.HashMap{
		"openid_list": openIDs,
		"tagid":       tagID,
	}
	_, err := comp.HttpPostJson("cgi-bin/tags/members/batchuntagging", params, nil, nil, result)

	return result, err
}
