package request

type RequestBroadcastModifyAssistant struct {
	RoomID int `json:"roomId"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
}
