package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type InterfaceSummary struct {
	RefDate       string `json:"ref_date"`
	CallbackCount int    `json:"callback_count"`
	FailCount     int    `json:"fail_count"`
	TotalTimeCost int    `json:"total_time_cost"`
	MaxTimeCost   int    `json:"max_time_cost"`
}

type ResponseDataCubeInterfaceSummary struct {
	*response.ResponseOfficialAccount

	List []*InterfaceSummary `json:"list"`
}

// ----------------------------------------------------------------------

type InterfaceSummaryHourly struct {
	RefDate       string `json:"ref_date"`
	RefHour       int    `json:"ref_hour"`
	CallbackCount int    `json:"callback_count"`
	FailCount     int    `json:"fail_count"`
	TotalTimeCost int    `json:"total_time_cost"`
	MaxTimeCost   int    `json:"max_time_cost"`
}

type ResponseDataCubeInterfaceSummaryHourly struct {
	*response.ResponseOfficialAccount

	List []*InterfaceSummaryHourly `json:"list"`
}
