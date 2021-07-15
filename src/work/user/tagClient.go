package user

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	response3 "github.com/ArtisanCloud/power-wechat/src/work/user/request"
	"github.com/ArtisanCloud/power-wechat/src/work/user/response"
)

type TagClient struct {
	*kernel.BaseClient
}

func NewTagClient(app kernel.ApplicationInterface) *TagClient {
	return &TagClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90209

// 创建标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90210
func (comp *TagClient) Create(tagName string, tagID int) *response.ResponseTagCreate {

	result := &response.ResponseTagCreate{}

	comp.HttpPostJson("cgi-bin/tag/create", &object.HashMap{
		"tagname": tagName,
		"tagid":   tagID,
	}, nil, nil, result)

	return result
}

// 更新标签名字
// https://open.work.weixin.qq.com/api/doc/90000/90135/90211
func (comp *TagClient) Update(data *response3.RequestUserDetail) *response2.ResponseWork {

	result := &response2.ResponseWork{}

	comp.HttpPostJson("cgi-bin/tag/update", data, nil, nil, result)

	return result
}

// 删除标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func (comp *TagClient) Delete(tagID string) *response2.ResponseWork {

	result := &response2.ResponseWork{}

	comp.HttpGet("cgi-bin/tag/delete", &object.StringMap{
		"tagid": tagID,
	}, nil, result)

	return result
}

// 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func (comp *TagClient) Get(tagID string) *response.ResponseTagDetail {

	result := &response.ResponseTagDetail{}

	comp.HttpGet("cgi-bin/tag/get", &object.StringMap{
		"tagid": tagID,
	}, nil, result)

	return result
}

// 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func (comp *TagClient) TagUsers(tagID int, userList []string) *response.ResponseTagCreateUser {
	return comp.tagOrUntagUsers("cgi-bin/tag/addtagusers", tagID, userList, []string{})
}

// 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func (comp *TagClient) TagDepartments(tagID int, partyList []string) *response.ResponseTagCreateUser {
	return comp.tagOrUntagUsers("cgi-bin/tag/addtagusers", tagID, []string{}, partyList)
}

func (comp *TagClient) tagOrUntagUsers(endpoint string, tagID int, userList []string, partyList []string) *response.ResponseTagCreateUser {

	result := &response.ResponseTagCreateUser{}

	comp.HttpPostJson(endpoint, &object.HashMap{
		"tagid":     tagID,
		"userlist":  userList,
		"partylist": partyList,
	}, nil, nil, result)

	return result
}

// 获取标签列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90216
func (comp *TagClient) List() *response.ResponseTagList {

	result := &response.ResponseTagList{}

	comp.HttpGet("cgi-bin/tag/list", nil, nil, result)

	return result
}
