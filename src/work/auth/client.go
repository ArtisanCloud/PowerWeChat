package auth

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/auth/response"
	userResponse "github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/response"
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

// 获取访问用户身份
// https://developer.work.weixin.qq.com/document/path/91023
func (comp *Client) GetUserInfo(ctx context.Context, code string) (*response.ResponseGetUserInfo, error) {

	result := new(response.ResponseGetUserInfo)
	req := object.StringMap{
		"code": code,
	}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/auth/getuserinfo", &req, nil, result)

	return result, err
}

// 获取访问用户敏感信息
// https://developer.work.weixin.qq.com/document/path/95833
func (comp *Client) GetUserDetail(ctx context.Context, userTicket string) (*userResponse.UserDetail, error) {

	result := new(userResponse.ResponseGetUserDetail)
	req := object.HashMap{
		"user_ticket": userTicket,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/auth/getuserdetail", &req, nil, nil, result)

	return result.UserDetail, err
}
