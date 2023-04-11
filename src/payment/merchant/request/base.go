package request

type Prepay interface {
	SetNotifyUrl(url string)
	SetSpMchId(id string)
	SetSubMchId(id string)
	SetSpAppId(id string)
	SetSubAppId(id string)

	GetNotifyUrl() string
	GetSpAppId() string
	GetSubAppId() string
	GetSpMchId() string
	GetSubMchId() string
}

type PrepayBase struct {
	SpMchId   string `json:"sp_mchid"`
	SubMchId  string `json:"sub_mchid"`
	SpAppId   string `json:"sp_appid"`
	SubAppId  string `json:"sub_appid"`
	NotifyUrl string `json:"notify_url"`
}

func (prepay *PrepayBase) SetNotifyUrl(url string) {
	prepay.NotifyUrl = url
}
func (prepay *PrepayBase) GetNotifyUrl() string {
	return prepay.NotifyUrl
}
func (prepay *PrepayBase) SetSpMchId(id string) {
	prepay.SpMchId = id
}
func (prepay *PrepayBase) SetSubMchId(id string) {
	prepay.SubMchId = id
}
func (prepay *PrepayBase) SetSpAppId(id string) {
	prepay.SpAppId = id
}
func (prepay *PrepayBase) SetSubAppId(id string) {
	prepay.SubAppId = id
}

func (prepay *PrepayBase) GetSpAppId() string {
	return prepay.SpAppId
}

func (prepay *PrepayBase) GetSubAppId() string {
	return prepay.SubAppId
}

func (prepay *PrepayBase) GetSpMchId() string {
	return prepay.SpMchId
}

func (prepay *PrepayBase) GetSubMchId() string {
	return prepay.SubMchId
}
