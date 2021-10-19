package linkedCorp

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取应用的可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93172
func (comp *Client) GetPermList() (*response.ResponseLinkCorpGetPermList, error) {

	result := &response.ResponseLinkCorpGetPermList{}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/agent/get_perm_list", nil, nil, nil, result)

	return result, err
}


// 获取互联企业成员详细信息
// https://work.weixin.qq.com/api/doc/90000/90135/93171
func (comp *Client) GetUser(userID string) (*response.ResponseLinkCorpGetUser, error) {

	result := &response.ResponseLinkCorpGetUser{}

	options := &object.HashMap{
		"userid":  userID,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/get", options, nil, nil, result)

	return result, err
}


// 获取互联企业部门成员
// https://work.weixin.qq.com/api/doc/90000/90135/93168
func (comp *Client) GetUserSimpleList(departmentID string, fetchChild bool) (*response.ResponseLinkCorpGetUserSimpleList, error) {

	result := &response.ResponseLinkCorpGetUserSimpleList{}

	options := &object.HashMap{
		"department_id":  departmentID,
		"fetch_child":  fetchChild,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/simplelist", options, nil, nil, result)

	return result, err
}


// 获取互联企业部门成员详情
// https://work.weixin.qq.com/api/doc/90000/90135/93169
func (comp *Client) GetUserList(departmentID string, fetchChild bool) (*response.ResponseLinkCorpGetUserList, error) {

	result := &response.ResponseLinkCorpGetUserList{}

	options := &object.HashMap{
		"department_id":  departmentID,
		"fetch_child":  fetchChild,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/list?", options, nil, nil, result)

	return result, err
}

// 获取互联企业部门列表
// https://work.weixin.qq.com/api/doc/90000/90135/93170
func (comp *Client) GetDepartmentList(departmentID string) (*response.ResponseLinkCorpGetDepartmentList, error) {

	result := &response.ResponseLinkCorpGetDepartmentList{}

	options := &object.HashMap{
		"department_id":  departmentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/department/list?", options, nil, nil, result)

	return result, err
}