package request

type RequestBroadcastGetSharedCode struct {
	RoomID int `json:"roomId"`
	Params string `json:"params"`
}
