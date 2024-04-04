package publish

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/publish/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/publish/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
	}, nil
}

// 新建草稿
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html
func (comp *Client) DraftAdd(ctx context.Context, data *request.RequestDraftAdd) (*response.ResponseDraftAdd, error) {
	result := &response.ResponseDraftAdd{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/add", data, &object.StringMap{}, nil, result)
	return result, err
}

// 获取草稿
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft.html
func (comp *Client) DraftGet(ctx context.Context, mediaID string) (*response.ResponseDraftGet, error) {
	result := &response.ResponseDraftGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/get", &object.HashMap{
		"media_id": mediaID,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 删除草稿
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Delete_draft.html
func (comp *Client) DraftDelete(ctx context.Context, mediaID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/delete", &object.HashMap{
		"media_id": mediaID,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 修改草稿
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Update_draft.html
func (comp *Client) DraftUpdate(ctx context.Context, data *request.RequestDraftUpdate) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/update", data, &object.StringMap{}, nil, result)

	return result, err
}

// 获取草稿总数
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Count_drafts.html
func (comp *Client) DraftCount(ctx context.Context) (*response.ResponseDraftCount, error) {
	result := &response.ResponseDraftCount{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/draft/count", &object.StringMap{}, nil, result)

	return result, err
}

// 获取草稿列表
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Count_drafts.html
func (comp *Client) DraftBatchGet(ctx context.Context, data *request.RequestBatchGet) (*response.ResponseBatchGet, error) {
	result := &response.ResponseBatchGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/batchget", data, &object.StringMap{}, nil, result)

	return result, err
}

// MP端开关
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Temporary_MP_Switch.html
func (comp *Client) DraftSwitch(ctx context.Context) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/switch", &object.HashMap{}, &object.StringMap{}, nil, result)

	return result, err
}

// M检查P端开关
// https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Temporary_MP_Switch.html
func (comp *Client) DraftCheckSwitch(ctx context.Context) (*response.ResponseCheckSwitch, error) {
	result := &response.ResponseCheckSwitch{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/draft/switch", &object.HashMap{}, &object.StringMap{
		"checkonly": "1",
	}, nil, result)

	return result, err
}

// 发布接口
// https://developers.weixin.qq.com/doc/offiaccount/Publish/Publish.html
func (comp *Client) PublishSubmit(ctx context.Context, mediaID string) (*response.ResponsePublishSubmit, error) {
	result := &response.ResponsePublishSubmit{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/freepublish/submit", &object.HashMap{
		"media_id": mediaID,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 发布状态轮询接口
// https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_status.html
func (comp *Client) PublishGet(ctx context.Context, publishID uint64) (*response.ResponsePublishGet, error) {
	result := &response.ResponsePublishGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/freepublish/get", &object.HashMap{
		"publish_id": publishID,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 删除发布
// https://developers.weixin.qq.com/doc/offiaccount/Publish/Delete_posts.html
func (comp *Client) PublishDelete(ctx context.Context, articleID string, index int) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/freepublish/delete", &object.HashMap{
		"article_id": articleID,
		"index":      index,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 通过 article_id 获取已发布文章
// https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_article_from_id.html
func (comp *Client) PublishGetArticle(ctx context.Context, articleID string) (*response.ResponsePublishGetArticle, error) {
	result := &response.ResponsePublishGetArticle{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/freepublish/getarticle", &object.HashMap{
		"article_id": articleID,
	}, &object.StringMap{}, nil, result)

	return result, err
}

// 获取成功发布列表
// https://developers.weixin.qq.com/doc/offiaccount/Publish/Get_publication_records.html
func (comp *Client) PublishBatchGet(ctx context.Context, data *request.RequestBatchGet) (*response.ResponseBatchGet, error) {
	result := &response.ResponseBatchGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/freepublish/batchget", data, &object.StringMap{}, nil, result)

	return result, err
}
