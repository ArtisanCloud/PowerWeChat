package corp

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suite "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	token := (*app).GetComponent("ProviderAccessToken").(*suite.AccessToken)
	baseClient, err := kernel.NewBaseClient(app, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

//// 获取代码草稿列表
//// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
//func (comp *Client) GetPreAuthorizationUrl() (*response.ResponseGetDrafts, error) {
//
//	result := &response.ResponseGetDrafts{}
//
//	_, err := comp.HttpGet("wxa/gettemplatedraftlist", nil, nil, result)
//
//	return result, err
//
//}
