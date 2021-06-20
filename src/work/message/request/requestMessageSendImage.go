package request

type RequestMessageSendImage struct {
	RequestMessageSend
	Image *RequestImage `json:"image"`
}

type RequestImage struct {
	MediaID string `json:"media_id"` //  "MEDIA_ID"
}
