package request

type ExternalContactList struct {
	TagList []string `json:"tag_list"`
}

type SenderList struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}

type VisibleRange struct {
	SenderList          SenderList          `json:"sender_list"`
	ExternalContactList ExternalContactList `json:"external_contact_list"`
}

type Text struct {
	Content string `json:"content"`
}

type Image struct {
	MediaId string `json:"media_id"`
}

type Link struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	MediaId string `json:"media_id"`
}

type Video struct {
	MediaId string `json:"media_id"`
}

type Attachment struct {
	Msgtype string `json:"msgtype"`
	Image   Image  `json:"image,omitempty"`
	Video   Video  `json:"video,omitempty"`
	Link    Link   `json:"link,omitempty"`
}

type RequestAddMomentTask struct {
	Text         Text          `json:"text"`
	Attachments  []Attachment  `json:"attachments"`
	VisibleRange *VisibleRange `json:"visible_range"`
}
