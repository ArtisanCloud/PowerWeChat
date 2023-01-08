package request

import "github.com/ArtisanCloud/PowerLibs/v3/object"

type TransferDetail struct {
	OutDetailNO    string            `json:"out_detail_no"`
	TransferAmount int               `json:"transfer_amount"`
	TransferRemark string            `json:"transfer_remark"`
	OpenID         string            `json:"openid"`
	UserName       object.NullString `json:"user_name"`
}

type RequestTransferBatch struct {
	AppID              string            `json:"appid"`
	OutBatchNO         string            `json:"out_batch_no"`
	BatchName          string            `json:"batch_name"`
	BatchRemark        string            `json:"batch_remark"`
	TotalAmount        int               `json:"total_amount"`
	TotalNum           int               `json:"total_num"`
	TransferDetailList []*TransferDetail `json:"transfer_detail_list"`
	TransferSceneID    string            `json:"transfer_scene_id,omitempty"`
}
