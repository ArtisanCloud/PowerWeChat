package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ButtonItem struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`

	NewsInfo struct {
		List []struct {
			Title      string `json:"title"`
			Author     string `json:"author"`
			Digest     string `json:"digest"`
			ShowCover  int    `json:"show_cover"`
			CoverUrl   string `json:"cover_url"`
			ContentUrl string `json:"content_url"`
			SourceUrl  string `json:"source_url"`
		} `json:"list"`
	} `json:"news_info"`
}

type SubButton struct {
	List []*ButtonItem `json:"list"`
}
type SelfButton struct {
	Type       string     `json:"type"`
	Name       string     `json:"name"`
	Key        string     `json:"key"`
	SubButtons *SubButton `json:"sub_button"`
}

// ----------------------------------------------------------------------

type SelfMenuInfo struct {
	Buttons []*SelfButton `json:"button"`
}

type ResponseCurrentSelfMenu struct {
	response.ResponseOfficialAccount

	IsMenuOpen   int           `json:"is_menu_open"`
	SelfMenuInfo *SelfMenuInfo `json:"selfmenu_info"`
}
