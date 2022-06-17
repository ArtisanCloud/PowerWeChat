package webdrive

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/webdrive/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/webdrive/response"
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
		baseClient,
	}, nil
}

// 新建空间
// https://work.weixin.qq.com/api/doc/90000/90135/93655
func (comp *Client) SpaceCreate(options *request.RequestWebDriveSpaceCreate) (*response.ResponseWebDriveSpaceCreate, error) {

	result := &response.ResponseWebDriveSpaceCreate{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_create", options, nil, nil, result)

	return result, err
}

// 重命名空间
// https://work.weixin.qq.com/api/doc/90000/90135/93655
func (comp *Client) SpaceRename(options *request.RequestWebDriveSpaceRename) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_rename", options, nil, nil, result)

	return result, err
}

// 解散空间
// https://work.weixin.qq.com/api/doc/90000/90135/93655
func (comp *Client) SpaceDismiss(options *request.RequestWebDriveSpaceDismiss) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_dismiss", options, nil, nil, result)

	return result, err
}

// 获取空间信息
// https://work.weixin.qq.com/api/doc/90000/90135/93655
func (comp *Client) SpaceInfo(options *request.RequestWebDriveSpaceInfo) (*response.ResponseWebDriveSpaceInfo, error) {

	result := &response.ResponseWebDriveSpaceInfo{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_info", options, nil, nil, result)

	return result, err
}

// 添加成员/部门
// https://work.weixin.qq.com/api/doc/90000/90135/93656
func (comp *Client) SpaceACLAdd(options *request.RequestWebDriveSpaceACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_acl_add", options, nil, nil, result)

	return result, err
}

// 移除成员/部门
// https://work.weixin.qq.com/api/doc/90000/90135/93656
func (comp *Client) SpaceACLDel(options *request.RequestWebDriveSpaceACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_acl_del", options, nil, nil, result)

	return result, err
}

// 权限管理
// https://work.weixin.qq.com/api/doc/90000/90135/93656
func (comp *Client) SpaceSetting(options *request.RequestWebDriveSpaceSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_setting", options, nil, nil, result)

	return result, err
}

// 获取邀请链接
// https://work.weixin.qq.com/api/doc/90000/90135/93656
func (comp *Client) SpaceShare(options *request.RequestWebDriveSpaceShare) (*response.ResponseWebDriveSpaceShare, error) {

	result := &response.ResponseWebDriveSpaceShare{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_share", options, nil, nil, result)

	return result, err
}

// 获取文件列表
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileList(options *request.RequestWebDriveFileList) (*response.ResponseWebDriveFileList, error) {

	result := &response.ResponseWebDriveFileList{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_list", options, nil, nil, result)

	return result, err
}

// 上传文件
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileUpload(options *request.RequestWebDriveFileUpload) (*response.ResponseWebDriveFileUpload, error) {

	result := &response.ResponseWebDriveFileUpload{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_upload", options, nil, nil, result)

	return result, err
}

// 下载文件
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileDownload(options *request.RequestWebDriveFileDownload) (*response.ResponseWebDriveFileDownload, error) {

	result := &response.ResponseWebDriveFileDownload{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_download", options, nil, nil, result)

	return result, err
}

// 新建文件/微文档
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileCreate(options *request.RequestWebDriveFileCreate) (*response.ResponseWebDriveFileCreate, error) {

	result := &response.ResponseWebDriveFileCreate{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_create", options, nil, nil, result)

	return result, err
}

// 重命名文件
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileRename(options *request.RequestWebDriveFileRename) (*response.ResponseWebDriveFileRename, error) {

	result := &response.ResponseWebDriveFileRename{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_rename", options, nil, nil, result)

	return result, err
}

// 移动文件
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileMove(options *request.RequestWebDriveFileMove) (*response.ResponseWebDriveFileMove, error) {

	result := &response.ResponseWebDriveFileMove{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_move", options, nil, nil, result)

	return result, err
}

// 删除文件
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func (comp *Client) FileDelete(options *request.RequestWebDriveFileDelete) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_delete", options, nil, nil, result)

	return result, err
}

// 新增指定人
// https://work.weixin.qq.com/api/doc/90000/90135/93658
func (comp *Client) FileACLAdd(options *request.RequestWebDriveFileACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_acl_add", options, nil, nil, result)

	return result, err
}

// 删除指定人
// https://work.weixin.qq.com/api/doc/90000/90135/93658
func (comp *Client) FileACLDel(options *request.RequestWebDriveFileACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_acl_del", options, nil, nil, result)

	return result, err
}

// 分享设置
// https://work.weixin.qq.com/api/doc/90000/90135/93658
func (comp *Client) FileSetting(options *request.RequestWebDriveFileSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_setting", options, nil, nil, result)

	return result, err
}

// 获取分享链接
// https://work.weixin.qq.com/api/doc/90000/90135/93658
func (comp *Client) FileShare(options *request.RequestWebDriveFileShare) (*response.ResponseWebDriveFileShare, error) {

	result := &response.ResponseWebDriveFileShare{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_share", options, nil, nil, result)

	return result, err
}
