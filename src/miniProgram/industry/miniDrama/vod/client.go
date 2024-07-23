package vod

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 拉取上传 该接口使用url传递到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-2-%E6%8B%89%E5%8F%96%E4%B8%8A%E4%BC%A0
func (comp *Client) VideoMediaUploadByURL(ctx context.Context, in *request.VideoMediaUploadByURLRequest) (result *response.VideoMediaUploadByURLResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}
	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/pullupload", params, nil, nil, &result)

	return
}

// 单文件上传
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-1-%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0
func (comp *Client) VideoMediaUploadByFile(ctx context.Context, in *request.VideoMediaUploadByFileRequest) (result *response.VideoMediaUploadByFileResponse, err error) {

	if in.MediaData == "" || in.MediaName == "" {
		return nil, errors.New("media_data is empty or media_name is empty")
	}

	files := &object.HashMap{
		"media_data": in.MediaData,
		"cover_data": in.CoverData,
	}

	if in.MediaType == "" {
		in.MediaType = "MP4"
	}

	if in.CoverType == "" {
		in.CoverType = "JPEG"
	}

	headerKV := &object.StringMap{
		"media_type": in.MediaType,
		"cover_type": in.CoverType,
		"media_name": in.MediaName,
	}
	ctx = context.WithValue(ctx, "headerKV", headerKV)

	_, err = comp.BaseClient.HttpUpload(ctx, "wxa/sec/vod/singlefileupload", files, nil, nil, nil, &result)

	return result, err
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

// 分片上传申请
// 上传大文件时需使用分片上传方式，分为 3 个步骤：
// 1.申请分片上传，确定文件名、格式类型，返回 upload_id，唯一标识本次分片上传。
// 2.上传分片，多次调用上传文件分片，需要携带 part_number 和 upload_id，其中 part_number 为分片的编号，支持乱序上传。当传入 part_number 和 upload_id 都相同的时候，后发起上传请求的分片将覆盖之前的分片。
// 3.确认分片上传，当上传完所有分片后，需要完成整个文件的合并。请求体中需要给出每一个分片的 part_number 和 etag，用来校验分片的准确性，最后返回文件的 media_id。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-4-%E7%94%B3%E8%AF%B7%E5%88%86%E7%89%87%E4%B8%8A%E4%BC%A0
func (comp *Client) ApplyMediaUploadId(ctx context.Context, in *request.VideoApplyChunkUploadByIdRequest) (result *response.VideoApplyChunkUploadByIdResponse, err error) {

	params, err := power.StructToHashMap(in)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/applyupload", params, nil, nil, &result)

	return
}

// 分片上传
// 将文件的其中一个分片上传到平台，最多支持100个分片，每个分片大小为5MB，最后一个分片可以小于5MB。该接口适用于视频和封面图片。视频最大支持500MB，封面图片最大支持10MB。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-5-%E4%B8%8A%E4%BC%A0%E5%88%86%E7%89%87
func (comp *Client) ApplyMediaChunkUpload(ctx context.Context, in *request.VideoApplyChunkUploadRequest) (result *response.VideoApplyChunkUploadResponse, err error) {

	if in.ChunkPath == "" {
		return nil, errors.New("chunk_path is empty")
	}

	fileParams := &object.HashMap{
		"data": in.ChunkPath,
	}

	headerKV := &object.HashMap{
		"resource_type": in.ResourceType,
		"part_number":   in.ChunkId,
		"upload_id":     in.UploadId,
	}
	ctx = context.WithValue(ctx, "headerKV", headerKV)

	_, err = comp.BaseClient.HttpUpload(ctx, "wxa/sec/vod/singlefileupload", fileParams, nil, nil, nil, &result)
	return
}

// 分片上传确认结束
// 该接口用于完成整个分片上传流程，合并所有文件分片，确认媒体文件（和封面图片文件）上传到平台的结果，返回文件的 ID。请求中需要给出每一个分片的 part_number 和 etag，用来校验分片的准确性。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_1-6-%E7%A1%AE%E8%AE%A4%E4%B8%8A%E4%BC%A0
func (comp *Client) UploadMediaChunkComplete(ctx context.Context, in *request.VideoChunkUploadCompleteRequest) (result *response.VideoChunkUploadCompleteResponse, err error) {

	params, err := power.StructToHashMap(in)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/commitupload", params, nil, nil, &result)
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
func (comp *Client) GetDramaList(ctx context.Context, in *request.ListRequest) (result *response.GetDramaListResponse, err error) {

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

// 提交替换剧集审核
// 该接口用于提交剧目替换剧集审核。待审核通过后7天内可使用替换审核通过的剧集接口（replacedramamedia）正式替换剧目中的相应的剧集，之后可以使用新的media_id播放。如7天内没有使用replacedramamedia进行剧集替换，则之后无法再发起替换。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-3-%E8%8E%B7%E5%8F%96%E5%89%A7%E7%9B%AE%E4%BF%A1%E6%81%AF
func (comp *Client) SubmitReplaceDramaAudit(ctx context.Context, in *request.SubmitReplaceDramaAuditRequest) (result *response.BaseResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/submitreplacedramamedias", params, nil, nil, &result)

	return
}

// 替换审核通过的剧集
// 该接口用于替换剧目审核通过的剧集。替换后可使用新的media_id获取播放链接，替换后旧media_id的播放链接不会马上失效，并且在7天内依然可以获取播放链接。为避免影响业务，调用该接口成功后，请及时替换获取播放链接时使用的media_id。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-5-%E6%9B%BF%E6%8D%A2%E5%AE%A1%E6%A0%B8%E9%80%9A%E8%BF%87%E7%9A%84%E5%89%A7%E9%9B%86
func (comp *Client) ReplaceAuditDrama(ctx context.Context, in *request.ReplaceAuditDramaRequest) (result *response.BaseResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/replacedramamedia", params, nil, nil, &result)

	return
}

// 修改剧目基本信息
// 该接口用于修改剧目基本信息。请求成功后，需要经过审核，审核通过后，最终才会修改基本信息。审核完成后，会下发通知。
// 1.剧目必须已经审核通过。
// 2.审核完成后会发送[事件通知]
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-6-%E4%BF%AE%E6%94%B9%E5%89%A7%E7%9B%AE%E5%9F%BA%E6%9C%AC%E4%BF%A1%E6%81%AF
func (comp *Client) UpdateDramaInfo(ctx context.Context, in *request.UpdateDramaInfoRequest) (result *response.BaseResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/modifydramabasicinfo", params, nil, nil, &result)

	return
}

// 查询剧目审核信息
// 该接口用于查询剧目最后一条审核信息。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_3-7-%E6%9F%A5%E8%AF%A2%E5%89%A7%E7%9B%AE%E5%AE%A1%E6%A0%B8%E4%BF%A1%E6%81%AF
func (comp *Client) SearchDramaAuditInfo(ctx context.Context, in *request.SearchDramaAuditInfoRequest) (result *response.SearchDramaAuditInfoResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getdramalatestauditinfo", params, nil, nil, &result)

	return
}

// 查询CDN用量数据
// 该接口用于查询点播 CDN 的流量数据。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_4-1-%E6%9F%A5%E8%AF%A2CDN%E7%94%A8%E9%87%8F%E6%95%B0%E6%8D%AE
func (comp *Client) SearchCdnUsageInfo(ctx context.Context, in *request.SearchCdnInfoRequest) (result *response.SearchCdnInfoResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getcdnusagedata", params, nil, nil, &result)

	return
}

// 查询CDN日志下载链接列表
// 查询域名的 CDN 访问日志的下载链接。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_4-2-%E6%9F%A5%E8%AF%A2CDN%E6%97%A5%E5%BF%97%E4%B8%8B%E8%BD%BD%E9%93%BE%E6%8E%A5%E5%88%97%E8%A1%A8
func (comp *Client) SearchCdnLogDownLoadUrl(ctx context.Context, in *request.SearchCdnInfoRequest) (result *response.SearchCdnLogDownLoadUrlResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getcdnlogs", params, nil, nil, &result)

	return
}

// 查询流量包详情
// 该接口用于查询流量包的使用详情。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_4-3-%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E5%8C%85%E8%AF%A6%E6%83%85
func (comp *Client) SearchListPackages(ctx context.Context, in *request.ListRequest) (result *response.ListPackageDetailResponse, err error) {

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

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/listpackages", params, nil, nil, &result)

	return
}

// 增加剧目授权
// 该接口用于授权方给被授权方授权一些剧目的播放权限。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_5-1-%E5%A2%9E%E5%8A%A0%E5%89%A7%E7%9B%AE%E6%8E%88%E6%9D%83
func (comp *Client) AddDramaAuthorized(ctx context.Context, in *request.DramaAuthorizedRequest) (result *response.DramaAuthorizedResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/authorizedrama", params, nil, nil, &result)

	return
}

// 解除剧目授权
// 该接口用于授权方解除被授权方一些剧目的播放授权。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_5-2-%E8%A7%A3%E9%99%A4%E5%89%A7%E7%9B%AE%E6%8E%88%E6%9D%83
func (comp *Client) RelieveDramaAuthorized(ctx context.Context, in *request.DramaAuthorizedRequest) (result *response.DramaAuthorizedResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/deauthorizedrama", params, nil, nil, &result)

	return
}

// 查询授权信息
// 该接口用于授权方查询剧目授权信息。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_5-3-%E6%9F%A5%E8%AF%A2%E6%8E%88%E6%9D%83%E4%BF%A1%E6%81%AF
func (comp *Client) SearchDramaAuthorized(ctx context.Context, in *request.SearchDramaAuthorizedRequest) (result *response.SearchDramaAuthorizedResponse, err error) {

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

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getauthorizeobjects", params, nil, nil, &result)

	return
}

// 查询被授权信息
// 该接口用于被授权方查询自己被授权的剧目授权信息
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_5-4-%E6%9F%A5%E8%AF%A2%E8%A2%AB%E6%8E%88%E6%9D%83%E4%BF%A1%E6%81%AF
func (comp *Client) SearchDoDramaAuthorized(ctx context.Context, in *request.SearchDoDramaAuthorizedRequest) (result *response.SearchDramaAuthorizedResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getauthorizedobjects", params, nil, nil, &result)

	return
}

// 增加账号授权
// 该接口用于授权方给被授权方授权所有剧目的播放权限。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_6-1-%E5%A2%9E%E5%8A%A0%E8%B4%A6%E5%8F%B7%E6%8E%88%E6%9D%83
func (comp *Client) AddAccountAuthorized(ctx context.Context, in *request.AccountAuthorizedRequest) (result *response.BaseResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/authorizeapp", params, nil, nil, &result)

	return
}

// 解除账号授权
// 该接口用于授权方解除被授权方的播放授权。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_6-2-%E8%A7%A3%E9%99%A4%E8%B4%A6%E5%8F%B7%E6%8E%88%E6%9D%83
func (comp *Client) RelieveAccountAuthorized(ctx context.Context, in *request.AccountAuthorizedRequest) (result *response.BaseResponse, err error) {

	params, err := power.StructToHashMap(in)

	if err != nil {

		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/deauthorizeapp", params, nil, nil, &result)

	return
}

// 查询账号授权信息
// 该接口用于授权方查询账号授权信息。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_6-3-%E6%9F%A5%E8%AF%A2%E8%B4%A6%E5%8F%B7%E6%8E%88%E6%9D%83%E4%BF%A1%E6%81%AF
func (comp *Client) SearchAccountAuthorized(ctx context.Context) (result *response.SearchAccountAuthorizedResponse, err error) {

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getauthorizeapps", nil, nil, nil, &result)

	return
}

// 查询被账号授权信息
// 该接口用于被授权方查询自己被账号授权的信息。
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/mini-drama/mini_drama.html#_6-4-%E6%9F%A5%E8%AF%A2%E8%A2%AB%E8%B4%A6%E5%8F%B7%E6%8E%88%E6%9D%83%E4%BF%A1%E6%81%AF
func (comp *Client) SearchDoAccountAuthorized(ctx context.Context) (result *response.SearchAccountAuthorizedResponse, err error) {

	_, err = comp.BaseClient.HttpPostJson(ctx, "wxa/sec/vod/getauthorizedapps", nil, nil, nil, &result)

	return
}
