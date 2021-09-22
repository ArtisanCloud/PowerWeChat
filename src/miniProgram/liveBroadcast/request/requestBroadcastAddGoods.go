package request

type RequestBroadcastAddGoods struct {
	IDs    []int `json:"ids"`
	RoomID int   `json:"roomId"`
}
