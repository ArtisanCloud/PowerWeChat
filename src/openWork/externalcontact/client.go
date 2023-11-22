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

// unionid查询external_userid
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

// 转换客户external_userid
// https://developer.work.weixin.qq.com/document/path/97063#%E8%BD%AC%E6%8D%A2%E5%AE%A2%E6%88%B7external-userid
func (comp *Client) GetNewExternalUserID(ctx context.Context, externalUserIDList []string) ([]response.NewExternalUserID, error) {

	result := new(response.ResponseGetNewExternalUserID)
	req := object.HashMap{
		"external_userid_list": externalUserIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_new_external_userid", &req, nil, nil, result)

	return result.Items, err
}

// 转换客户群成员external_userid
// https://developer.work.weixin.qq.com/document/path/97063#%E8%BD%AC%E6%8D%A2%E5%AE%A2%E6%88%B7%E7%BE%A4%E6%88%90%E5%91%98external-userid
func (comp *Client) GetGroupNewExternalUserID(ctx context.Context, chatID string, externalUserIDList []string) ([]response.NewExternalUserID, error) {

	result := new(response.ResponseGetNewExternalUserID)
	req := object.HashMap{
		"chatid":               chatID,
		"external_userid_list": externalUserIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/groupchat/get_new_external_userid", &req, nil, nil, result)

	return result.Items, err
}
