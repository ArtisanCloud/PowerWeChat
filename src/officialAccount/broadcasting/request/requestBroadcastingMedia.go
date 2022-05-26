package request

type MediaPreviewInterface interface {
	GetMsgType() string
	GetToUser() string
}

// ---------------------------------------------
type MediaPreview struct {
	MediaPreviewInterface

	ToUser   string `json:"touser"`
	ToWXName string `json:"towxname"`
	MsgType  string `json:"msgtype"`
}

func (media *MediaPreview) GetMsgType() string {
	return media.MsgType
}

func (media *MediaPreview) GetToUser() string {
	return media.ToUser
}

// ---------------------------------------------

type MediaInfo struct {
	MediaID string `json:"media_id"`
}

type NPNews struct {
	*MediaPreview

	MPNews *MediaInfo `json:"mpnews"`
}

type Text struct {
	*MediaPreview

	Text MediaInfo `json:"text"`
}

type Voice struct {
	*MediaPreview

	Voice MediaInfo `json:"voice"`
}

type Image struct {
	*MediaPreview

	Image MediaInfo `json:"image"`
}

type MPVideo struct {
	*MediaPreview

	MPVideo MediaInfo `json:"mpvideo"`
}

type CardExt struct {
	Code      string `json:"code"`
	Openid    string `json:"openid "`
	Timestamp string `json:"timestamp "`
	Signature string `json:"signature "`
}

type WXCardInfo struct {
	CardID  string   `json:"card_id"`
	CardExt *CardExt `json:"card_ext"`
}

type WXCard struct {
	*MediaPreview

	WXCard *WXCardInfo `json:"wxcard"`
}
