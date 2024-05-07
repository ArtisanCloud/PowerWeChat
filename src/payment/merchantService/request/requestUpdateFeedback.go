package request

type RequestUpdateFeedback struct {
	Action          string   `json:"action"`
	LaunchRefundDay int      `json:"launch_refund_day"`
	RejectReason    string   `json:"reject_reason"`
	RejectMediaList []string `json:"reject_media_list"`
	Remark          string   `json:"remark"`
}
