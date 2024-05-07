package response

import "time"

type Data struct {
	LogId              string             `json:"log_id"`
	Operator           string             `json:"operator"`
	OperateTime        time.Time          `json:"operate_time"`
	OperateType        string             `json:"operate_type"`
	OperateDetails     string             `json:"operate_details"`
	ImageList          []string           `json:"image_list"`
	ComplaintMediaList ComplaintMediaList `json:"complaint_media_list"`
}

type ResponseQueryNegotiationHistoriesByComplaint struct {
	Data       []Data `json:"data"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	TotalCount int    `json:"total_count"`
}
