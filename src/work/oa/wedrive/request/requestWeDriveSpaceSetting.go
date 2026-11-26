package request

type RequestWeDriveSpaceSetting struct {
	UserID                       string `json:"userid"`
	SpaceID                      string `json:"spaceid"`
	EnableWatermark              bool   `json:"enable_watermark"`
	AddMemberOnlyAdmin           bool   `json:"add_member_only_admin"`
	EnableShareUrl               bool   `json:"enable_share_url"`
	ShareUrlNoApprove            bool   `json:"share_url_no_approve"`
	ShareUrlNoApproveDefaultAuth int    `json:"share_url_no_approve_default_auth"`
}

// RequestWeDriveSpaceSettingPatch only changes supplied settings. A nil field leaves the existing setting in place.
type RequestWeDriveSpaceSettingPatch struct {
	SpaceID                      string `json:"spaceid"`
	EnableWatermark              *bool  `json:"enable_watermark,omitempty"`
	EnableConfidentialMode       *bool  `json:"enable_confidential_mode,omitempty"`
	ShareUrlNoApprove            *bool  `json:"share_url_no_approve,omitempty"`
	ShareUrlNoApproveDefaultAuth *int   `json:"share_url_no_approve_default_auth,omitempty"`
	DefaultFileScope             *int   `json:"default_file_scope,omitempty"`
	BanShareExternal             *bool  `json:"ban_share_external,omitempty"`
}
