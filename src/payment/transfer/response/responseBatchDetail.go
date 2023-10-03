package response

import "time"

type TransferBatchDetail struct {
	MchId          string    `json:"mchid"`
	OutBatchNo     string    `json:"out_batch_no"`
	BatchId        string    `json:"batch_id"`
	Appid          string    `json:"appid"`
	OutDetailNo    string    `json:"out_detail_no"`
	DetailId       string    `json:"detail_id"`
	DetailStatus   string    `json:"detail_status"`
	TransferAmount int       `json:"transfer_amount"`
	TransferRemark string    `json:"transfer_remark"`
	FailReason     string    `json:"fail_reason"`
	Openid         string    `json:"openid"`
	UserName       string    `json:"user_name"`
	InitiateTime   time.Time `json:"initiate_time"`
	UpdateTime     time.Time `json:"update_time"`
}
