package request

type RequestLivingModify struct {
	LivingID       string `json:"livingid"`
	Theme          string `json:"theme"`
	LivingStart    int    `json:"living_start"`
	LivingDuration int    `json:"living_duration"`
	Description    string `json:"description"`
	Type           int    `json:"type"`
	RemindTime     int    `json:"remind_time"`
}
