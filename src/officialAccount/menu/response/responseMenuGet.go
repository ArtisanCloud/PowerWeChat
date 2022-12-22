package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type SubButton struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
	Key      string `json:"key,omitempty"`
}
type Button struct {
	Type       string       `json:"type"`
	Name       string       `json:"name"`
	Key        string       `json:"key"`
	SubButtons []*SubButton `json:"sub_button"`
}

type Menu struct {
	Buttons []*Button `json:"button"`
	MenuID  int       `json:"menuid"`
}

type MatchRule struct {
	GroupID            int    `json:"group_id"`
	Sex                int    `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType int    `json:"client_platform_type"`
}

type ConditionalMenu struct {
	Buttons   []*Button  `json:"button"`
	MatchRule *MatchRule `json:"matchrule"`
	MenuID    int        `json:"menuid"`
}

type ResponseMenuGet struct {
	*response.ResponseOfficialAccount

	Menus            *Menu              `json:"menu"`
	ConditionalMenus []*ConditionalMenu `json:"conditionalmenu"`
}

// ----------------------------------------------------------------------

type SelfMenuInfo struct {
	Buttons []*Button `json:"button"`
}

type ResponseCurrentSelfMenu struct {
	*response.ResponseOfficialAccount

	IsMenuOpen   int           `json:"is_menu_open"`
	SelfMenuInfo *SelfMenuInfo `json:"selfmenu_info"`
}
