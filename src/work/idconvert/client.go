package idconvert

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/idconvert/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// unionid转换为第三方external_userid
// https://developer.work.weixin.qq.com/document/path/95900#unionid%E8%BD%AC%E6%8D%A2%E4%B8%BA%E7%AC%AC%E4%B8%89%E6%96%B9external-userid
func (comp *Client) UnionIDToExternalUserID(ctx context.Context, unionID string, openID string, subjectType int) (*response.ResponseUnionIDToExternalUserID, error) {

	result := new(response.ResponseUnionIDToExternalUserID)
	req := object.HashMap{
		"unionid": unionID,
		"openid":  openID,
	}
	if subjectType == 1 {
		req["subject_type"] = subjectType
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/idconvert/unionid_to_external_userid", &req, nil, nil, result)

	return result, err
}

// external_userid查询pending_id
// 该接口可用于当一个微信用户成为企业客户前已经使用过服务商服务（服务商曾通过unionid查询external_userid接口获取到pending_id）的场景。本接口获取到的pending_id可以维持unionid和external_userid的关联关系。pending_id有效期为90天，超过有效期之后，将无法通过该接口将external_userid换取对应的pending_id。
func (comp *Client) ExternalUserIDToPendingID(ctx context.Context, externalUserID []string, chatID string) ([]response.ExternalUserIDToPendingIDResult, error) {

	result := new(response.ResponseExternalUserIDToPendingID)
	req := object.HashMap{
		"external_userid": externalUserID,
	}
	if chatID != "" {
		req["chat_id"] = chatID
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/idconvert/batch/external_userid_to_pending_id", &req, nil, nil, result)

	return result.Result, err
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

// userid的转换
// https://developer.work.weixin.qq.com/document/path/97062
func (comp *Client) OpenUserIDToUserID(ctx context.Context, sourceAgentID int, openUserIDList []string) (*response.ResponseOpenUserIDToUserID, error) {

	result := new(response.ResponseOpenUserIDToUserID)
	req := object.HashMap{
		"source_agentid":   sourceAgentID,
		"open_userid_list": openUserIDList,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/openuserid_to_userid", &req, nil, nil, result)

	return result, err
}

// 客户标签ID的转换
// https://developer.work.weixin.qq.com/document/path/96169
func (comp *Client) ExternalTagIDToOpenExternalTagID(ctx context.Context, externalTagIDList []string) (*response.ResponseExternalTagIDToOpenExternalTagID, error) {

	result := new(response.ResponseExternalTagIDToOpenExternalTagID)
	req := object.HashMap{
		"external_tagid_list": externalTagIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/idconvert/external_tagid", &req, nil, nil, result)

	return result, err
}
