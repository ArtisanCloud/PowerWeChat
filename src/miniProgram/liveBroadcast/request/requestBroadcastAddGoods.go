package request

type RequestBroadcastAddGoods struct {
	IDs string `json:"ids"`
	RoomID int `json:"roomId"`
}
