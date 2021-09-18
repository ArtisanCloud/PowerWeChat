package request

type RequestResignedTransferResult struct {
	HandoverUserID string `json:"handover_userid" `
	TakeoverUserID string `json:"takeover_userid"`
	Cursor         string `json:"cursor"`
}
