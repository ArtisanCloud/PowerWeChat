package request

type RequestMenuSet struct {
	Button []RequestMenuSetButton `json:"button"`
}

type RequestMenuSetButton struct {
	Type      string                 `json:"type"`
	Name      string                 `json:"name"`
	Key       string                 `json:"key,omitempty"`      // 用于普通点击菜单
	PagePath  string                 `json:"pagepath,omitempty"` // 小程序菜单
	AppID     string                 `json:"appid,omitempty"`    // 小程序appid
	Url       string                 `json:"url,omitempty"`      // 网页链接，成员点击菜单可打开链接
	SubButton []RequestMenuSetButton `json:"sub_button"`
}
