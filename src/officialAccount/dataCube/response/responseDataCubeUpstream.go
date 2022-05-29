package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type UpstreamMessageSummary struct {
	RefDate  string `json:"ref_date"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

type ResponseDataCubeUpstreamMessageSummary struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageSummary `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageHourly struct {
	RefDate  string `json:"ref_date"`
	RefHour  int    `json:"ref_hour"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

type ResponseDataCubeUpstreamMessageHourly struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageHourly `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageWeekly struct {
	RefDate  string `json:"ref_date"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

type ResponseDataCubeUpstreamMessageWeekly struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageWeekly `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageMonthly struct {
	RefDate  string `json:"ref_date"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

type ResponseDataCubeUpstreamMessageMonthly struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageMonthly `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageDistSummary struct {
	RefDate       string `json:"ref_date"`
	CountInterval int    `json:"count_interval"`
	MsgUser       int    `json:"msg_user"`
}

type ResponseDataCubeUpstreamMessageDistSummary struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageDistSummary `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageDistWeekly struct {
	RefDate       string `json:"ref_date"`
	CountInterval int    `json:"count_interval"`
	MsgUser       int    `json:"msg_user"`
}

type ResponseDataCubeUpstreamMessageDistWeekly struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageDistWeekly `json:"list"`
}

// ----------------------------------------------------------------------

type UpstreamMessageDistMonthly struct {
	RefDate       string `json:"ref_date"`
	CountInterval int    `json:"count_interval"`
	MsgUser       int    `json:"msg_user"`
}

type ResponseDataCubeUpstreamMessageDistMonthly struct {
	*response.ResponseOfficialAccount

	List []*UpstreamMessageDistMonthly `json:"list"`
}

// ----------------------------------------------------------------------
