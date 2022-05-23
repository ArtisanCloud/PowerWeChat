package wifi

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/user/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取用户基本信息
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (comp *Client) Get(openID string, lang string) (*response.ResponseGetUserInfo, error) {

	result := &response.ResponseGetUserInfo{}

	params := &object.StringMap{
		"openid": openID,
		"lang":   lang,
	}

	_, err := comp.HttpGet("cgi-bin/user/info", params, nil, result)

	return result, err

}
