package request

type RequestUserDetail struct {
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	Mobile         string `json:"mobile"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Enable         int    `json:"enable"`
	AvatarMediaid  string `json:"avatar_mediaid"`
	Telephone      string `json:"telephone"`
	Address        string `json:"address"`
	MainDepartment int    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	ToInvite         bool   `json:"to_invite"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}