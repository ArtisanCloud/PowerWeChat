package request

type RequestWeDriveSpaceSetting struct {
	SpaceID                      string `json:"spaceid"`                                     // 空间id 必填
	EnableWatermark              bool   `json:"enable_watermark,omitempty"`                  // （本字段仅专业版企业可设置）启用水印。false:关 true:开 ;如果不填充此字段为保持原有状态
	EnableConfidentialMode       bool   `json:"enable_confidential_mode,omitempty"`          // 是否开启保密模式。false:关 true:开 如果不填充此字段为保持原有状态
	ShareUrlNoApprove            bool   `json:"share_url_no_approve,omitempty"`              // 通过链接加入空间无需审批。false:关； true:开； 如果不填充此字段为保持原有状态
	ShareUrlNoApproveDefaultAuth int    `json:"share_url_no_approve_default_auth,omitempty"` // 邀请链接默认权限。1:仅下载 2:可编辑 4:仅预览 5:可上传下载 200:自定义权限；如果不填充此字段为保持原有状态
	DefaultFileScope             int    `json:"default_file_scope,omitempty"`                // 文件默认可查看范围。1:仅成员；2:企业内。如果不填充此字段为保持原有状态
	BanShareExternal             bool   `json:"ban_share_external,omitempty"`                // 是否禁止文件分享到企业外｜false:关 true:开 如果不填充此字段为保持原有状态
}
