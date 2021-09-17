package request

type RequestBroadcastGetLiveReplay struct {
	Action string `json:"action"`
	RoomID int `json:"room_id"`
	Start  int `json:"start"`
	Limit  int `json:"limit"`
}
