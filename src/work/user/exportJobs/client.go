package exportJobs

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/user/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 导出成员
// https://work.weixin.qq.com/api/doc/90000/90135/94849
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
// https://work.weixin.qq.com/api/doc/90000/90135/94851
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
// https://work.weixin.qq.com/api/doc/90000/90135/94852
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
// https://work.weixin.qq.com/api/doc/90000/90135/94853
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
// https://work.weixin.qq.com/api/doc/90000/90135/94854
func (comp *Client) GetExportResult(jobID string) (*response.ResponseUserExportGetResult, error) {

	result := &response.ResponseUserExportGetResult{}

	options := &object.HashMap{
		"jobid": jobID,
	}
	_, err := comp.HttpPostJson("cgi-bin/export/get_result", options, nil, nil, result)

	return result, err
}
