package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type DataSummary struct {
	ClickPV int `json:"click_pv"`
	ClickUV int `json:"click_uv"`
	FTime   int `json:"ftime"`
	ShakePV int `json:"shake_pv"`
	ShakeUV int `json:"shake_uv"`
}

type ResponseStatsSummary struct {
	*response.ResponseOfficialAccount

	Data []DataSummary `json:"data"`
}

// ------------------------------------------------ {

type StateDevice struct {
	DeviceID int    `json:"device_id"`
	Major    string `json:"major"`
	Minor    string `json:"minor"`
	UUID     string `json:"uuid"`
	ClickPV  int    `json:"click_pv"`
	ClickUV  int    `json:"click_uv"`
	ShakePV  int    `json:"shake_pv"`
	ShakeUV  int    `json:"shake_uv"`
}

type StatsDeviceList struct {
	Devices []*StateDevice `json:"devices"`
}

type ResponseStatsDeviceList struct {
	*response.ResponseOfficialAccount

	Data       *StatsDeviceList `json:"data"`
	Date       int              `json:"date"`
	TotalCount int              `json:"total_count"`
	PageIndex  int              `json:"page_index"`
	Success    string           `json:"success."`
}

// ------------------------------------------------

type ResponseStatsPage struct {
	*response.ResponseOfficialAccount

	Data []*DataSummary `json:"data"`
}

// ------------------------------------------------

type StatsPage struct {
	PageID  int `json:"page_id"`
	ClickPV int `json:"click_pv"`
	ClickUV int `json:"click_uv"`
	ShakePV int `json:"shake_pv"`
	ShakeUV int `json:"shake_uv"`
}

type StatePageList struct {
	Pages []StatsPage `json:"pages"`
}

type ResponseStatsPageList struct {
	*response.ResponseOfficialAccount

	Data       *StatePageList `json:"data"`
	Date       int            `json:"date"`
	TotalCount int            `json:"total_count"`
	PageIndex  int            `json:"page_index"`
}
