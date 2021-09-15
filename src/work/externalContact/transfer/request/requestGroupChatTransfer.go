package request

type RequestGroupChatTransfer struct {
	ChatIDList []string `json:"chat_id_list" `
	NewOwner   string   `json:"new_owner"`
}
