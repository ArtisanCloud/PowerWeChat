package exportJobs

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/response"
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

// 导出成员
// https://developer.work.weixin.qq.com/document/path/94849
func (comp *Client) SimpleUser(encodingAESKey string, blockSize int64) (*response.ResponseUserExportJobs, error) {

	result := &response.ResponseUserExportJobs{}

	options := &object.HashMap{
		"encoding_aeskey": encodingAESKey,
		"block_size":      blockSize,
	}
	_, err := comp.HttpPostJson("cgi-bin/export/simple_user", options, nil, nil, result)

	return result, err
}

// 导出成员详情
// https://developer.work.weixin.qq.com/document/path/94851
func (comp *Client) User(encodingAESKey string, blockSize int64) (*response.ResponseUserExportJobs, error) {

	result := &response.ResponseUserExportJobs{}

	options := &object.HashMap{
		"encoding_aeskey": encodingAESKey,
		"block_size":      blockSize,
	}
	_, err := comp.HttpPostJson("cgi-bin/export/user", options, nil, nil, result)

	return result, err
}

// 导出部门
// https://developer.work.weixin.qq.com/document/path/94852
func (comp *Client) Department(encodingAESKey string, blockSize int64) (*response.ResponseUserExportJobs, error) {

	result := &response.ResponseUserExportJobs{}

	options := &object.HashMap{
		"encoding_aeskey": encodingAESKey,
		"block_size":      blockSize,
	}
	_, err := comp.HttpPostJson("cgi-bin/export/department", options, nil, nil, result)

	return result, err
}

// 导出标签成员
// https://developer.work.weixin.qq.com/document/path/94853
func (comp *Client) TagUser(tagID int, encodingAESKey string, blockSize int64) (*response.ResponseUserExportJobs, error) {

	result := &response.ResponseUserExportJobs{}

	options := &object.HashMap{
		"tagid":           tagID,
		"encoding_aeskey": encodingAESKey,
		"block_size":      blockSize,
	}
	_, err := comp.HttpPostJson("cgi-bin/export/taguser", options, nil, nil, result)

	return result, err
}

// 获取导出结果
// https://developer.work.weixin.qq.com/document/path/94854
func (comp *Client) GetExportResult(jobID string) (*response.ResponseUserExportGetResult, error) {

	result := &response.ResponseUserExportGetResult{}

	options := &object.StringMap{
		"jobid": jobID,
	}
	_, err := comp.HttpGet("cgi-bin/export/get_result", options, nil, result)

	return result, err
}
