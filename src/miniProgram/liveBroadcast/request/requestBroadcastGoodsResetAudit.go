package request

type RequestBroadcastGoodsResetAudit struct {
	AuditID int `json:"auditId"`
	GoodsID int `json:"goodsId"`
}
