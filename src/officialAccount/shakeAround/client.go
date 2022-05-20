package shakeAround

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/response"
)

type Client struct {
	*kernel.BaseClient
}

// 申请开通摇一摇周边功能
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Apply/Application_for_opening_function.html
func (comp *Client) Register(data *request.RequestShakeAroundAccountRegister) (*response.ResponseShakeAroundAccountRegister, error) {

	result := &response.ResponseShakeAroundAccountRegister{}

	_, err := comp.HttpPostJson("shakearound/account/register", data, nil, nil, result)

	return result, err
}

// 查询审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Apply/Application_for_opening_function.html
func (comp *Client) Status() (*response.ResponseShakeAroundAccountRegister, error) {

	result := &response.ResponseShakeAroundAccountRegister{}

	_, err := comp.HttpGet("shakearound/account/auditstatus", nil, nil, result)

	return result, err
}

// 获取设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Getting_Device_and_User_Information.html
func (comp *Client) User(data *request.RequestShakeAroundUser) (*response.ResponseShakeAroundUser, error) {

	result := &response.ResponseShakeAroundUser{}

	_, err := comp.HttpPostJson("shakearound/user/getshakeinfo", data, nil, nil, result)

	return result, err
}

func (comp *Client) UserWithPoi(ticket string) (*response.ResponseShakeAroundUser, error) {

	data := &request.RequestShakeAroundUser{
		Ticket:  ticket,
		NeedPoi: 1,
	}

	return comp.User(data)
}
