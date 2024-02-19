package externalcontact

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/externalcontact/response"
	suite "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
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

// 第三方主体unionid转换为第三方external_userid
// https://developer.work.weixin.qq.com/document/path/96516#42-unionid%E6%9F%A5%E8%AF%A2external-userid
func (comp *Client) UnionIDToExternalUserID3rd(ctx context.Context, unionID string, openID string, corpID string) ([]response.ExternalUserIDInfo, error) {

	var result response.UnionIDToExternalUserID3rd
	req := object.StringMap{
		"unionid": unionID,
		"openid":  openID,
	}
	if corpID != "" {
		req["corpid"] = corpID
	}
	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/externalcontact/unionid_to_external_userid_3rd", &req, nil, &result)

	return result.ExternalUserIDInfo, err
}
