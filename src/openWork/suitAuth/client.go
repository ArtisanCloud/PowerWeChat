package suit

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	// token := app.GetComponent("SuiteAccessToken").(*AccessToken)
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取预授权码
// https://developer.work.weixin.qq.com/document/10975#%E8%8E%B7%E5%8F%96%E9%A2%84%E6%8E%88%E6%9D%83%E7%A0%81
func (comp *Client) GetPreAuthCode(ctx context.Context) (*response.PreAuthCode, error) {

	var result response.PreAuthCode

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/service/get_pre_auth_code", nil, nil, &result)

	return &result, err
}

// 设置授权配置
// https://developer.work.weixin.qq.com/document/10975#%E8%AE%BE%E7%BD%AE%E6%8E%88%E6%9D%83%E9%85%8D%E7%BD%AE
func (comp *Client) SetSessionInfo(ctx context.Context, req *request.SetSessionInfoRequest) error {
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/service/set_session_info", req, nil, nil, nil)
	return err
}

// 获取企业永久授权码
// https://developer.work.weixin.qq.com/document/10975#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%B0%B8%E4%B9%85%E6%8E%88%E6%9D%83%E7%A0%81
func (comp *Client) GetPermanentCode(ctx context.Context, authCode string) (*response.GetPermanentCodeResponse, error) {

	var result response.GetPermanentCodeResponse
	req := object.HashMap{
		"auth_code": authCode,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/service/get_permanent_code", &req, nil, nil, &result)

	return &result, err
}

// 获取企业授权信息
// https://developer.work.weixin.qq.com/document/10975#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%8E%88%E6%9D%83%E4%BF%A1%E6%81%AF
func (comp *Client) GetAuthInfo(ctx context.Context, authCorpID string, permanentCode string) (*response.GetPermanentCodeResponse, error) {

	var result response.GetPermanentCodeResponse
	req := object.HashMap{
		"auth_corp_id":   authCorpID,
		"permanent_code": permanentCode,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/service/get_auth_info", &req, nil, nil, &result)

	return &result, err
}

// userid的转换
// https://developer.work.weixin.qq.com/document/path/97062
func (comp *Client) UserIDToOpenUserID(ctx context.Context, userIDList []string) ([]response.UserIDToOpenUserIDResult, error) {

	result := new(response.ResponseUserIDToOpenUserID)
	req := object.HashMap{
		"userid_list": userIDList,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/userid_to_openuserid", &req, nil, nil, result)

	return result.OpenUserIDList, err
}
