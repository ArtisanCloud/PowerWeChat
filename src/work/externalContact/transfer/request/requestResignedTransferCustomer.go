package request

type RequestResignedTransferCustomer struct {
	HandoverUserID string   `json:"handover_userid" `
	TakeoverUserID string   `json:"takeover_userid"`
	ExternalUserID []string `json:"external_userid"`
}
