package request

type RequestAppChatCreate struct {

	Name string `json:"name"`
	Owner string `json:"owner"`
	UserList []string `json:"userlist"`
	ChatID string `json:"chatid"`

}
