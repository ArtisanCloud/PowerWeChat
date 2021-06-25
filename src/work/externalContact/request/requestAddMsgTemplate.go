package request

type RequestAddMsgTemplate struct {
	ExternalUserID []string      `json:"external_userid" binding:"required"`
	Sender         string        `json:"sender" binding:"required"`
	Text           TextOfMessage `json:"text" binding:"required"`
	Attachments    []interface{} `json:"attachments"`
}
