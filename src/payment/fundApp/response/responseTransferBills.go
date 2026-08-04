package response

import (
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

// https://work.weixin.qq.com/api/doc/90000/90135/90278

type ResponseTransferBills struct {
	response.ResponsePayment
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	CreateTime     string `json:"create_time"`
	State          string `json:"state"`
	FailReason     string `json:"fail_reason"`
	PackageInfo    string `json:"package_info"`
}

type ResponseCancelBill struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	State          string `json:"state"`
	UpdateTime     string `json:"update_time"`
}

type ResponseQueryOutBill struct {
	response.ResponsePayment
	MchId          string `json:"mch_id"`
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	Appid          string `json:"appid"`
	State          string `json:"state"`
	TransferAmount int    `json:"transfer_amount"`
	TransferRemark string `json:"transfer_remark"`
	FailReason     string `json:"fail_reason"`
	Openid         string `json:"openid"`
	UserName       string `json:"user_name"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type ResponseQueryTransferBill struct {
	MchId          string    `json:"mch_id"`
	OutBillNo      string    `json:"out_bill_no"`
	TransferBillNo string    `json:"transfer_bill_no"`
	Appid          string    `json:"appid"`
	State          string    `json:"state"`
	TransferAmount int       `json:"transfer_amount"`
	TransferRemark string    `json:"transfer_remark"`
	FailReason     string    `json:"fail_reason"`
	Openid         string    `json:"openid"`
	UserName       string    `json:"user_name"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}
