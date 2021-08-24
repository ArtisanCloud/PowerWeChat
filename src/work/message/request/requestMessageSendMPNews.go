package request

type RequestMessageSendMPNews struct {
	RequestMessageSend
	MPNews struct {
		Article []*RequestMPNewsArticle `json:"articles"`
	} `json:"mpnews"`
}

type RequestMPNewsArticle struct {
	Title            string `json:"title"`              // "Title",
	ThumbMediaID     string `json:"thumb_media_id"`     // "MEDIA_ID",
	Author           string `json:"author"`             // "Author",
	ContentSourceURL string `json:"content_source_url"` // "URL",
	Content          string `json:"content"`            // "Content",
	Digest           string `json:"digest"`             // "Digest description"
}
