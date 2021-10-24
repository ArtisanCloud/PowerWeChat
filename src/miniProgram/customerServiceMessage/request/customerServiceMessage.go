package request

type CustomerServiceMsgText struct {
	Content string `json:"content"`
}

type CustomerServiceMsgImage struct {
	MediaID string `json:"media_id"`
}

type CustomerServiceMsgLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumb_url"`
}

type CustomerServiceMsgMpPage struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}
