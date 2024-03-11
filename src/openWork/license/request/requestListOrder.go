package request

type RequestListOrder struct {
	CorpID string `json:"corpid,omitempty"`
	// StartTime 开始时间,下单时间。可不填。但是不能单独指定该字段，start_time跟end_time必须同时指定。
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间,下单时间。起始时间跟结束时间不能超过31天。可不填。但是不能单独指定该字段，start_time跟end_time必须同时指定。
	EndTime string `json:"end_time,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500
	Limit int `json:"limit,omitempty"`
}
