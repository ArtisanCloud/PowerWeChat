package user

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suite "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/user/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	token := (*app).GetComponent("SuiteAccessToken").(*suite.AccessToken)
	baseClient, err := kernel.NewBaseClient(app, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 第三方根据code获取企业成员信息
// https://developer.work.weixin.qq.com/document/10975#%E7%AC%AC%E4%B8%89%E6%96%B9%E6%A0%B9%E6%8D%AEcode%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF
func (comp *Client) GetUserInfo3rdByCode(ctx context.Context, code string) (*response.GetUserInfo3rdByCode, error) {

	var result response.GetUserInfo3rdByCode
	req := object.StringMap{
		"code": code,
	}
	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/service/getuserinfo3rd", &req, nil, &result)

	return &result, err
}

// 第三方使用user_ticket获取成员详情
// https://developer.work.weixin.qq.com/document/10975#%E7%AC%AC%E4%B8%89%E6%96%B9%E4%BD%BF%E7%94%A8user-ticket%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E8%AF%A6%E6%83%85
func (comp *Client) GetUserInfo3rdByUserTicket(ctx context.Context, userTicket string) (*response.GetUserInfo3rdByUserTicket, error) {
	var result response.GetUserInfo3rdByUserTicket
	req := object.StringMap{
		"user_ticket": userTicket,
	}
	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/service/getuserinfo3rd", &req, nil, &result)
	return &result, err
}
