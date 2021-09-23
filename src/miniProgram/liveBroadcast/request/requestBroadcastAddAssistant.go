package request

type RequestBroadcastAddAssistantUser struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type RequestBroadcastAddAssistant struct {
	RoomID int                                `json:"roomId"`
	Users  []RequestBroadcastAddAssistantUser `json:"users"`
}
