package request

type RequestWebDriveSpaceSetting struct {
	UserID                       string `json:"userid"`
	SpaceID                      string `json:"spaceid"`
	EnableWatermark              bool   `json:"enable_watermark"`
	AddMemberOnlyAdmin           bool   `json:"add_member_only_admin"`
	EnableShareUrl               bool   `json:"enable_share_url"`
	ShareUrlNoApprove            bool   `json:"share_url_no_approve"`
	ShareUrlNoApproveDefaultAuth int    `json:"share_url_no_approve_default_auth"`
}
