package request

type Prepay interface {
	SetNotifyUrl(url string)
	SetAppID(id string)
	SetMchID(id string)
	GetNotifyUrl() string
	GetAppID() string
}

type PrepayBase struct {
	AppID     string
	MchID     string
	NotifyUrl string
}

func (prepay *PrepayBase) SetNotifyUrl(url string) {
	prepay.NotifyUrl = url
}
func (prepay *PrepayBase) GetNotifyUrl() string {
	return prepay.NotifyUrl
}
func (prepay *PrepayBase) SetMchID(id string) {
	prepay.MchID = id
}
func (prepay *PrepayBase) SetAppID(id string) {
	prepay.AppID = id
}
func (prepay *PrepayBase) GetAppID() string {
	return prepay.AppID
}
