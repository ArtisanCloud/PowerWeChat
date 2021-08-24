package request

type RequestMessageSendVideo struct {
	RequestMessageSend
	Video *RequestVideo `json:"video"`
}

type RequestVideo struct {
	MediaID     string `json:"media_id"`    //  "MEDIA_ID",
	Title       string `json:"title"`       // "Title",
	Description string `json:"description"` // "Description"
}
