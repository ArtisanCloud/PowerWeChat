package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Category struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	FirstID     int    `json:"first_id"`
	SecondID    int    `json:"second_id"`
	ThirdClass  string `json:"third_class,omitempty"`
	ThirdID     int    `json:"third_id,omitempty"`
}

type ResponseGetCategory struct {
	response.ResponseOpenPlatform

	CategoryList []*Category `json:"category_list"`
}

type ResponseGetPage struct {
	response.ResponseOpenPlatform

	PageList []string `json:"page_list"`
}

type ResponseSubmitAudit struct {
	response.ResponseOpenPlatform

	Type    string `json:"type"`
	MediaID string `json:"mediaid"`
}

type ResponseGetAuditStatus struct {
	response.ResponseOpenPlatform

	Status     int    `json:"status"`
	Reason     string `json:"reason"`
	Screenshot string `json:"screenshot"`
}

type ResponseGetLatestAuditStatus struct {
	response.ResponseOpenPlatform

	AuditID         int    `json:"auditid"`
	Status          int    `json:"status"`
	Reason          string `json:"reason"`
	ScreenShot      string `json:"ScreenShot"`
	UserVersion     string `json:"user_version"`
	UserDesc        string `json:"user_desc"`
	SubmitAuditTime int    `json:"submit_audit_time"`
}

type Version struct {
	CommitTime  int    `json:"commit_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	AppVersion  int    `json:"app_version"`
}

type ResponseRollbackRelease struct {
	response.ResponseOpenPlatform

	VersionList []Version `json:"version_list"`
}

type GrayReleasePlan struct {
	Status                  int  `json:"status"`
	CreateTimestamp         int  `json:"create_timestamp"`
	GrayPercentage          int  `json:"gray_percentage"`
	SupportExperiencerFirst bool `json:"support_experiencer_first"`
	SupportDebugerFirst     bool `json:"support_debuger_first"`
}

type ResponseGetGrayRelease struct {
	response.ResponseOpenPlatform

	GrayReleasePlan *GrayReleasePlan `json:"gray_release_plan"`
}

// --------------------------------------------------------------------------------
type Item struct {
	Percentage int    `json:"percentage"`
	Version    string `json:"version"`
}

type UVInfo struct {
	Items []Item `json:"items"`
}

type ResponseGetSupportVersion struct {
	response.ResponseOpenPlatform

	NowVersion string  `json:"now_version"`
	UVInfo     *UVInfo `json:"uv_info"`
}

// --------------------------------------------------------------------------------

type ResponseQueryQuota struct {
	response.ResponseOpenPlatform

	Rest         int `json:"rest"`
	Limit        int `json:"limit"`
	SpeedupRest  int `json:"speedup_rest"`
	SpeedupLimit int `json:"speedup_limit"`
}
