package request

type SubButton struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
	Key      string `json:"key,omitempty"`
}

type Button struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name"`
	Key       string `json:"key,omitempty"`
	SubButtons []SubButton `json:"sub_button,omitempty"`
}

type RequestMenuCreate struct {
	Buttons []*Button `json:"button"`
}

