package request

type RequestBroadcastRemoveAssistant struct {
	RoomID int `json:"roomId"`
	UserName string `json:"username"`
}
