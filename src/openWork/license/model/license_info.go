package model

type LicenseInfo struct {
	// LicenseStatus license检查开启状态。
	// 0：未开启license检查状态（未迁移的历史授权的第三方应用（接入版本付费）或者未达到拦截时间的历史授权的的第三方应用（未接入版本付费）以及代开发应用）
	// 1：已开启license检查状态。若开启且已过试用期，则需要为企业购买license账号才可以使用
	LicenseStatus int `json:"license_status,omitempty"`
	// LicenseCheckTime 接口开启拦截校验时间。开始拦截校验后，无接口许可将会被拦截，有接口许可将不会被拦截。
	LicenseCheckTime int64 `json:"license_check_time,omitempty"`
	// TrailInfo 应用license试用期信息。仅当license_status为1且应用有试用期时返回该字段。服务商测试企业、历史迁移应用无试用期。
	TrailInfo *TrailInfo `json:"trail_info,omitempty"`
}

type TrailInfo struct {
	// StartTime 接口许可试用开始时间
	StartTime int64 `json:"start_time,omitempty"`
	// EndTime 接口许可试用到期时间。若企业多次安装卸载同一个第三方应用，以第一次安装的时间为试用期开始时间，第一次安装完90天后为结束试用时间。
	EndTime int64 `json:"end_time,omitempty"`
}
