package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"time"
)

type ResponseTrasferBatch struct {
	response.ResponsePayment

	OutBatchNo string    `json:"out_batch_no"`
	BatchId    string    `json:"batch_id"`
	CreateTime time.Time `json:"create_time"`
}

type TransferBatch struct {
	MchID         string    `json:"mchid"`
	OutBatchNO    string    `json:"out_batch_no"`
	BatchID       string    `json:"batch_id"`
	AppID         string    `json:"appid"`
	BatchStatus   string    `json:"batch_status"`
	BatchType     string    `json:"batch_type"`
	BatchName     string    `json:"batch_name"`
	BatchRemark   string    `json:"batch_remark"`
	CloseReason   string    `json:"close_reason"`
	TotalAmount   int       `json:"total_amount"`
	TotalNum      int       `json:"total_num"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	SuccessAmount int       `json:"success_amount"`
	SuccessNum    int       `json:"success_num"`
	FailAmount    int       `json:"fail_amount"`
	FailNum       int       `json:"fail_num"`
}

type TransferDetail struct {
	DetailID     string `json:"detail_id"`
	OutDetailNO  string `json:"out_detail_no"`
	DetailStatus string `json:"detail_status"`
}

type ResponseTrasferQueryBatch struct {
	response.ResponsePayment

	TransferBatch      *TransferBatch    `json:"transfer_batch"`
	TransferDetailList []*TransferDetail `json:"transfer_detail_list"`
}

type ResponseTrasferQueryOutBatchNO struct {
	response.ResponsePayment

	Limit              int               `json:"limit"`
	Offset             int               `json:"offset"`
	TransferBatch      *TransferBatch    `json:"transfer_batch"`
	TransferDetailList []*TransferDetail `json:"transfer_detail_list"`
}

type ResponseOutBatchNODetail struct {
	response.ResponsePayment

	OutBatchNO     string    `json:"out_batch_no"`
	BatchID        string    `json:"batch_id"`
	AppID          string    `json:"appid"`
	OutDetailNO    string    `json:"out_detail_no"`
	DetailID       string    `json:"detail_id"`
	DetailStatus   string    `json:"detail_status"`
	TransferAmount int       `json:"transfer_amount"`
	TransferRemark string    `json:"transfer_remark"`
	FailReason     string    `json:"fail_reason"`
	OpenID         string    `json:"openid"`
	UserName       string    `json:"user_name"`
	InitiateTime   time.Time `json:"initiate_time"`
	UpdateTime     time.Time `json:"update_time"`
}

type ResponseTrasferBillReceipt struct {
	response.ResponsePayment

	OutBatchNO      string    `json:"out_batch_no"`
	SignatureNO     string    `json:"signature_no"`
	SignatureStatus string    `json:"signature_status"`
	HashType        string    `json:"hash_type"`
	HashValue       string    `json:"hash_value"`
	DownloadURL     string    `json:"download_url"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

type ResponseTrasferElectronicReceipts struct {
	AcceptType      string `json:"accept_type"`
	OutBatchNO      string `json:"out_batch_no"`
	OutDetailNO     string `json:"out_detail_no"`
	SignatureNO     string `json:"signature_no"`
	SignatureStatus string `json:"signature_status"`
	HashType        string `json:"hash_type"`
	HashValue       string `json:"hash_value"`
	DownloadURL     string `json:"download_url"`
}
