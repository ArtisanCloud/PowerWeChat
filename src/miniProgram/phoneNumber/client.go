package phoneNumber

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/phoneNumber/response"
)

type Client struct {
	*kernel.BaseClient
}

// code换取用户手机号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (comp *Client) GetUserPhoneNumber(code string) (*response.ResponseGetUserPhoneNumber, error) {

	result := &response.ResponseGetUserPhoneNumber{}
	_, err := comp.HttpPostJson("wxa/business/getuserphonenumber", &object.HashMap{
		"code": code,
	}, nil, nil, result)

	return result, err
}
