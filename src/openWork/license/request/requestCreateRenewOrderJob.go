package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"

type RequestCreateRenewOrderJob struct {
	CorpID string `json:"corpid,omitempty"`
	// AccountList 续期的账号列表，每次最多1000个。同一个jobid最多关联1000000个基础账号跟1000000个互通账号
	AccountList []model.ActiveInfo `json:"account_list,omitempty"`
	// JobID 任务id，若不传则默认创建一个新任务。若指定第一次调用后拿到jobid，可以通过该接口将jobid关联多个userid
	JobID string `json:"jobid,omitempty"`
}
