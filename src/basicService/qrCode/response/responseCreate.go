package response

type ResponseQRCodeCreate struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds "`
	Url           string `json:"url "`
}