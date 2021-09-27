package request

type RequestAppChatUpdate struct {

	ChatID string `json:"chatid"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`


}
