package tag

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/response"
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

// 获取企业标签库
// https://developer.work.weixin.qq.com/document/path/92117
func (comp *Client) GetCorpTagList(ctx context.Context, tagID []string, groupID []string) (*response.ResponseTagGetCorpTagList, error) {

	result := &response.ResponseTagGetCorpTagList{}

	options := &object.HashMap{
		"tag_id":   tagID,
		"group_id": groupID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_corp_tag_list", options, nil, nil, result)

	return result, err
}

// 添加企业客户标签
// https://developer.work.weixin.qq.com/document/path/92117
func (comp *Client) AddCorpTag(ctx context.Context, options *request.RequestTagAddCorpTag) (*response.ResponseTagAddCorpTag, error) {

	result := &response.ResponseTagAddCorpTag{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/add_corp_tag", options, nil, nil, result)

	return result, err
}

// 编辑企业客户标签
// https://developer.work.weixin.qq.com/document/path/92117
func (comp *Client) EditCorpTag(ctx context.Context, options *request.RequestTagEditCorpTag) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/edit_corp_tag", options, nil, nil, result)

	return result, err
}

// 删除企业客户标签
// https://developer.work.weixin.qq.com/document/path/92117
func (comp *Client) DelCorpTag(ctx context.Context, options *request.RequestTagDelCorpTag) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/del_corp_tag", options, nil, nil, result)

	return result, err
}

// 获取指定规则组下的企业客户标签
// https://developer.work.weixin.qq.com/document/path/94882
func (comp *Client) GetStrategyTagList(ctx context.Context, options *request.RequestTagGetStrategyTagList) (*response.ResponseTagGetStrategyTagList, error) {

	result := &response.ResponseTagGetStrategyTagList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_strategy_tag_list", options, nil, nil, result)

	return result, err
}

// 为指定规则组创建企业客户标签
// https://developer.work.weixin.qq.com/document/path/94882
func (comp *Client) AddStrategyTag(ctx context.Context, options *request.RequestTagAddStrategyTag) (*response.ResponseTagAddStrategyTag, error) {

	result := &response.ResponseTagAddStrategyTag{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/add_strategy_tag", options, nil, nil, result)

	return result, err
}

// 编辑指定规则组下的企业客户标签
// https://developer.work.weixin.qq.com/document/path/94882
func (comp *Client) EditStrategyTag(ctx context.Context, options *request.RequestTagEditStrategyTag) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/edit_strategy_tag", options, nil, nil, result)

	return result, err
}

// 删除指定规则组下的企业客户标签
// https://developer.work.weixin.qq.com/document/path/94882
func (comp *Client) DelStrategyTag(ctx context.Context, options *request.RequestTagDelStrategyTag) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/edit_strategy_tag", options, nil, nil, result)

	return result, err
}

// 编辑客户企业标签
// https://developer.work.weixin.qq.com/document/path/92118
func (comp *Client) MarkTag(ctx context.Context, options *request.RequestTagMarkTag) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/mark_tag", options, nil, nil, result)

	return result, err
}
