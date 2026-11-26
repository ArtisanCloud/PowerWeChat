package request

type GroupList struct {
	TagList []string `json:"tag_list"`
}

type TagFilter struct {
	GroupList []GroupList `json:"group_list"`
}

type RequestAddMsgTemplate struct {
	ChatType       string                     `json:"chat_type"`
	ExternalUserID []string                   `json:"external_userid"`
	ChatIdList     []string                   `json:"chat_id_list"`
	TagFilter      TagFilter                  `json:"tag_filter,omitempty"`
	Sender         string                     `json:"sender,omitempty"`
	AllowSelect    bool                       `json:"allow_select"`
	Text           *TextOfMessage             `json:"text,omitempty"`
	Attachments    []MessageTemplateInterface `json:"attachments,omitempty"`
}
