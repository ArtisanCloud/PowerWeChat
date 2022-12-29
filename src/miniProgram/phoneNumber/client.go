package phoneNumber

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/phoneNumber/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// code换取用户手机号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (comp *Client) GetUserPhoneNumber(ctx context.Context, code string) (*response.ResponseGetUserPhoneNumber, error) {

	result := &response.ResponseGetUserPhoneNumber{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/getuserphonenumber", &object.HashMap{
		"code": code,
	}, nil, nil, result)

	return result, err
}
