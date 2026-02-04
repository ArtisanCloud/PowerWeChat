package wedrive

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/wedrive/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/wedrive/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 新建空间
// https://developer.work.weixin.qq.com/document/path/93655#新建空间
func (comp *Client) SpaceCreate(ctx context.Context, options *request.RequestWeDriveSpaceCreate) (*response.ResponseWeDriveSpaceCreate, error) {

	result := &response.ResponseWeDriveSpaceCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_create", options, nil, nil, result)

	return result, err
}

// 重命名空间
// https://developer.work.weixin.qq.com/document/path/93655#重命名空间
func (comp *Client) SpaceRename(ctx context.Context, options *request.RequestWeDriveSpaceRename) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_rename", options, nil, nil, result)

	return result, err
}

// 解散空间
// https://developer.work.weixin.qq.com/document/path/93655#解散空间
func (comp *Client) SpaceDismiss(ctx context.Context, options *request.RequestWeDriveSpaceDismiss) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_dismiss", options, nil, nil, result)

	return result, err
}

// 获取空间信息
// https://developer.work.weixin.qq.com/document/path/93655#获取空间相册信息
func (comp *Client) SpaceInfo(ctx context.Context, options *request.RequestWeDriveSpaceInfo) (*response.ResponseWeDriveSpaceInfo, error) {

	result := &response.ResponseWeDriveSpaceInfo{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_info", options, nil, nil, result)

	return result, err
}

// 添加成员/部门
// https://developer.work.weixin.qq.com/document/path/93656#添加成员部门
func (comp *Client) SpaceACLAdd(ctx context.Context, options *request.RequestWeDriveSpaceACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_acl_add", options, nil, nil, result)

	return result, err
}

// 移除成员/部门
// https://developer.work.weixin.qq.com/document/path/93656#移除成员部门
func (comp *Client) SpaceACLDel(ctx context.Context, options *request.RequestWeDriveSpaceACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_acl_del", options, nil, nil, result)

	return result, err
}

// 安全设置
// https://developer.work.weixin.qq.com/document/path/93656#安全设置
func (comp *Client) SpaceSetting(ctx context.Context, options *request.RequestWeDriveSpaceSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_setting", options, nil, nil, result)

	return result, err
}

// 获取邀请链接
// https://developer.work.weixin.qq.com/document/path/93656#获取邀请链接
func (comp *Client) SpaceShare(ctx context.Context, options *request.RequestWeDriveSpaceShare) (*response.ResponseWeDriveSpaceShare, error) {

	result := &response.ResponseWeDriveSpaceShare{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/space_share", options, nil, nil, result)

	return result, err
}

// 获取文件列表
// https://developer.work.weixin.qq.com/document/path/93657#获取文件列表
func (comp *Client) FileList(ctx context.Context, options *request.RequestWeDriveFileList) (*response.ResponseWeDriveFileList, error) {

	result := &response.ResponseWeDriveFileList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_list", options, nil, nil, result)

	return result, err
}

// 上传文件
// https://developer.work.weixin.qq.com/document/path/93657#上传文件
func (comp *Client) FileUpload(ctx context.Context, options *request.RequestWeDriveFileUpload) (*response.ResponseWeDriveFileUpload, error) {

	result := &response.ResponseWeDriveFileUpload{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_upload", options, nil, nil, result)

	return result, err
}

// 下载文件
// https://developer.work.weixin.qq.com/document/path/93657#下载文件
func (comp *Client) FileDownload(ctx context.Context, options *request.RequestWeDriveFileDownload) (*response.ResponseWeDriveFileDownload, error) {

	result := &response.ResponseWeDriveFileDownload{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_download", options, nil, nil, result)

	return result, err
}

// 分块上传初始化
// https://developer.work.weixin.qq.com/document/path/98004#分块上传初始化
func (comp *Client) FileUploadInit(ctx context.Context, options *request.RequestWeDriveFileUploadInit) (*response.ResponseWeDriveFileUploadInit, error) {

	result := &response.ResponseWeDriveFileUploadInit{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_upload_init", options, nil, nil, result)

	return result, err
}

// 分块上传文件
// https://developer.work.weixin.qq.com/document/path/98004#分块上传文件
func (comp *Client) FileUploadPart(ctx context.Context, options *request.RequestWeDriveFileUploadPart) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_upload_part", options, nil, nil, result)

	return result, err
}

// 分块上传完成
// https://developer.work.weixin.qq.com/document/path/98004#分块上传完成
func (comp *Client) FileUploadFinish(ctx context.Context, options *request.RequestWeDriveFileUploadFinish) (*response.ResponseWeDriveFileUploadFinish, error) {

	result := &response.ResponseWeDriveFileUploadFinish{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_upload_finish", options, nil, nil, result)

	return result, err
}

// 新建文件/微文档
// https://developer.work.weixin.qq.com/document/path/93657#新建文件文档
func (comp *Client) FileCreate(ctx context.Context, options *request.RequestWeDriveFileCreate) (*response.ResponseWeDriveFileCreate, error) {

	result := &response.ResponseWeDriveFileCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_create", options, nil, nil, result)

	return result, err
}

// 重命名文件
// https://developer.work.weixin.qq.com/document/path/93657#重命名文件
func (comp *Client) FileRename(ctx context.Context, options *request.RequestWeDriveFileRename) (*response.ResponseWeDriveFileRename, error) {

	result := &response.ResponseWeDriveFileRename{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_rename", options, nil, nil, result)

	return result, err
}

// 移动文件
// https://developer.work.weixin.qq.com/document/path/93657#移动文件
func (comp *Client) FileMove(ctx context.Context, options *request.RequestWeDriveFileMove) (*response.ResponseWeDriveFileMove, error) {

	result := &response.ResponseWeDriveFileMove{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_move", options, nil, nil, result)

	return result, err
}

// 删除文件
// https://developer.work.weixin.qq.com/document/path/93657#删除文件
func (comp *Client) FileDelete(ctx context.Context, options *request.RequestWeDriveFileDelete) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_delete", options, nil, nil, result)

	return result, err
}

// 新增成员
// https://developer.work.weixin.qq.com/document/path/93658
func (comp *Client) FileACLAdd(ctx context.Context, options *request.RequestWeDriveFileACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_acl_add", options, nil, nil, result)

	return result, err
}

// 删除成员
// https://developer.work.weixin.qq.com/document/path/93658#删除成员
func (comp *Client) FileACLDel(ctx context.Context, options *request.RequestWeDriveFileACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_acl_del", options, nil, nil, result)

	return result, err
}

// 分享设置
// https://developer.work.weixin.qq.com/document/path/93658#分享设置
func (comp *Client) FileSetting(ctx context.Context, options *request.RequestWeDriveFileSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_setting", options, nil, nil, result)

	return result, err
}

// 获取分享链接
// https://developer.work.weixin.qq.com/document/path/95860#获取分享链接
func (comp *Client) FileShare(ctx context.Context, options *request.RequestWeDriveFileShare) (*response.ResponseWeDriveFileShare, error) {

	result := &response.ResponseWeDriveFileShare{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedrive/file_share", options, nil, nil, result)

	return result, err
}
