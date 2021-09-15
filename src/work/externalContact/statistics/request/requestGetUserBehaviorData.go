package request

type RequestGetUserBehaviorData struct {
	UserID    []string `json:"userid"`
	PartyID   []int    `json:"partyid"`
	StartTime int64    `json:"start_time" `
	EndTime   int64    `json:"end_time"`
}
