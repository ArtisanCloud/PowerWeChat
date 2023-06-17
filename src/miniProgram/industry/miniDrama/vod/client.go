package vod

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 拉取上传 该接口使用url传递到微信服务器
//https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-2-%E6%8B%89%E5%8F%96%E4%B8%8A%E4%BC%A0

func (comp *Client) VideoMediaUploadByURL(ctx context.Context, in *request.VideoMediaUploadByURLRequest) (result *response.VideoMediaUploadByURLResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}
	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/pullupload", params, nil, nil, &result)

	return
}

// 查询任务 通过上传接口获得的taskid 进行查询
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-2-%E6%8B%89%E5%8F%96%E4%B8%8A%E4%BC%A0

func (comp *Client) SearchMediaByTaskId(ctx context.Context, taskId int64) (result *response.VideoMediaSearchTaskResponse, err error) {

	params := &power.HashMap{
		"task_id": taskId,
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/gettask", params, nil, nil, &result)

	return

}

// 获取媒资列表 该接口用于获取已上传到平台的媒资列表。
// 本接口返回的视频或图片链接均为临时链接，不应将其保存下来。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-1-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E5%88%97%E8%A1%A8

func (comp *Client) GetMediaList(ctx context.Context, in *request.MediaListRequest) (result *response.MediaListResponse, err error) {

	if in.Limit > 100 {
		in.Limit = 100
	}

	if in.Offset <= 0 {
		in.Offset = 0
	}

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/listmedia", params, nil, nil, &result)

	return

}

// 获取媒资详细信息
// 该接口用于获取已上传到平台的指定媒资信息，用于开发者后台管理使用。用于给用户客户端播放的链接应该使用getmedialink接口获取
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-2-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF

func (comp *Client) GetMediaInfo(ctx context.Context, mediaId int64) (result *response.MediaInfoResponse, err error) {

	params := &power.HashMap{
		"media_id": mediaId,
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getmedia", params, nil, nil, &result)

	return

}

// 获取媒资播放链接
// 需要提审后才能调用该接口，不然会返回未审核通过
// 该接口用于获取视频临时播放链接，用于给用户的播放使用。只有审核通过的视频才能通过该接口获取播放链接。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-3-%E8%8E%B7%E5%8F%96%E5%AA%92%E8%B5%84%E6%92%AD%E6%94%BE%E9%93%BE%E6%8E%A5

func (comp *Client) GetMediaLink(ctx context.Context, in *request.GetMediaLinkRequest) (result *response.MediaLinkResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getmedialink", params, nil, nil, &result)

	return

}

// 删除媒资
// 该接口用于删除指定媒资。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_2-4-%E5%88%A0%E9%99%A4%E5%AA%92%E8%B5%84

func (comp *Client) DeleteMedia(ctx context.Context, mediaId int64) (result *response.BaseResponse, err error) {

	params := &power.HashMap{
		"media_id": mediaId,
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/deletemedia", params, nil, nil, &result)

	return
}

// 剧目提审
// 该接口用于提交剧目审核。
// 剧目信息与审核材料在首次提审时为必填，重新提审时根据是否需要修改选填，
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-1-%E5%89%A7%E7%9B%AE%E6%8F%90%E5%AE%A1

func (comp *Client) SubmitAudit(ctx context.Context, in *request.SubmitAuditRequest) (result *response.SubmitAuditResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/auditdrama", params, nil, nil, &result)

	return
}

// 获取剧目列表
// 该接口用于获取已提交的剧目列表。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-2-%E8%8E%B7%E5%8F%96%E5%89%A7%E7%9B%AE%E5%88%97%E8%A1%A8

func (comp *Client) GetDramaList(ctx context.Context, in *request.GetDramaListRequest) (result *response.GetDramaListResponse, err error) {

	if in.Limit > 100 {
		in.Limit = 100
	}

	if in.Offset < 0 {
		in.Offset = 0
	}

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/listdramas", params, nil, nil, &result)

	return
}

// 获取剧目信息
// 该接口用于查询已提交的剧目。
// 本接口返回的图片链接均为临时链接，不应将其保存下来。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-3-%E8%8E%B7%E5%8F%96%E5%89%A7%E7%9B%AE%E4%BF%A1%E6%81%AF

func (comp *Client) GetDramaInfo(ctx context.Context, dramaId int64) (result *response.GetDramaInfoResponse, err error) {

	params := &power.HashMap{
		"drama_id": dramaId,
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getdrama", params, nil, nil, &result)

	return
}
