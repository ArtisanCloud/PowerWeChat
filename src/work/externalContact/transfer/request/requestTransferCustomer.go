package request

type RequestTransferCustomer struct {
	HandoverUserID     string   `json:"handover_userid" `
	TakeoverUserID     string   `json:"takeover_userid"`
	ExternalUserID     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg"`
}
