package request

type RequestGetPerformanceData struct {
	Time   *GetPerformanceDataTime     `json:"time"`
	Module string                      `json:"module"`
	Params []*GetPerformanceDataParams `json:"params"`
}

type GetPerformanceDataTime struct {
	// int64类型： time.Now().Unix()
	BeginTimestamp int64 `json:"begin_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`
}

type GetPerformanceDataParams struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
