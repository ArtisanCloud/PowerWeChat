package tag

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/response"
)

type Client struct {
	*kernel.BaseClient
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
// https://developer.work.weixin.qq.com/document/path/90210
func (comp *Client) Create(tagName string, tagID int64) (*response.ResponseTagCreate, error) {

	result := &response.ResponseTagCreate{}

	_, err := comp.HttpPostJson("cgi-bin/tag/create", &object.HashMap{
		"tagname": tagName,
		"tagid":   tagID,
	}, nil, nil, result)

	return result, err
}

// 更新标签名字
// https://developer.work.weixin.qq.com/document/path/90211
func (comp *Client) Update(tagName string, tagID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/tag/update", &object.HashMap{
		"tagname": tagName,
		"tagid":   tagID,
	}, nil, nil, result)

	return result, err
}

// 删除标签
// https://developer.work.weixin.qq.com/document/path/90212
func (comp *Client) Delete(tagID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpGet("cgi-bin/tag/delete", &object.StringMap{
		"tagid": fmt.Sprintf("%d", tagID),
	}, nil, result)

	return result, err
}

// 获取标签成员
// https://developer.work.weixin.qq.com/document/path/90832
func (comp *Client) Get(tagID int64) (*response.ResponseTagDetail, error) {

	result := &response.ResponseTagDetail{}

	_, err := comp.HttpGet("cgi-bin/tag/get", &object.StringMap{
		"tagid": fmt.Sprintf("%d", tagID),
	}, nil, result)

	return result, err
}

// 获取标签成员
// https://developer.work.weixin.qq.com/document/path/90833
func (comp *Client) TagUsers(tagID int64, userList []string) (*response.ResponseTagCreateUser, error) {
	return comp.TagOrUntagUsers("cgi-bin/tag/addtagusers", tagID, userList, []string{})
}

// 增加标签成员
// https://developer.work.weixin.qq.com/document/path/90833
func (comp *Client) TagDepartments(tagID int64, partyList []string) (*response.ResponseTagCreateUser, error) {
	return comp.TagOrUntagUsers("cgi-bin/tag/addtagusers", tagID, []string{}, partyList)
}

func (comp *Client) TagOrUntagUsers(endpoint string, tagID int64, userList []string, partyList []string) (*response.ResponseTagCreateUser, error) {

	result := &response.ResponseTagCreateUser{}

	_, err := comp.HttpPostJson(endpoint, &object.HashMap{
		"tagid":     tagID,
		"userlist":  userList,
		"partylist": partyList,
	}, nil, nil, result)

	return result, err
}

// 获取标签列表
// https://developer.work.weixin.qq.com/document/path/90216
func (comp *Client) List() (*response.ResponseTagList, error) {

	result := &response.ResponseTagList{}

	_, err := comp.HttpGet("cgi-bin/tag/list", nil, nil, result)

	return result, err
}
