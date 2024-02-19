package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseCreateRenewOrderJob struct {
	response.ResponseWork
	// JobID 任务id，请求包中未指定jobid时，会生成一个新的jobid返回
	JobID string `json:"jobid,omitempty"`
	// InvalidAccountList 不合法的续期账号列表
	InvalidAccountList []InvalidAccount `json:"invalid_account_list,omitempty"`
}

type InvalidAccount struct {
	response.ResponseWork
	model.ActiveInfo
}
