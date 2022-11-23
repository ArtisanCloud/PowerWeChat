package request

type RequestUserDetail struct {
	Userid           string          `json:"userid"`
	Name             string          `json:"name"`
	Alias            string          `json:"alias"`
	Mobile           string          `json:"mobile"`
	Department       []int           `json:"department"`
	Order            []int           `json:"order"`
	Position         string          `json:"position"`
	Gender           uint32          `json:"gender"`
	Email            string          `json:"email"`
	BizMail          string          `json:"biz_mail"`
	IsLeaderInDept   []int           `json:"is_leader_in_dept"`
	DirectLeader     []string        `json:"direct_leader"`
	Enable           int             `json:"enable"`
	AvatarMediaid    string          `json:"avatar_mediaid"`
	Telephone        string          `json:"telephone"`
	Address          string          `json:"address"`
	MainDepartment   int             `json:"main_department"`
	Extattr          Extattr         `json:"extattr"`
	ToInvite         bool            `json:"to_invite"`
	ExternalPosition string          `json:"external_position"`
	ExternalProfile  ExternalProfile `json:"external_profile"`
}
type Text struct {
	Value string `json:"value"`
}
type Web struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
type Attrs struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text Text   `json:"text,omitempty"`
	Web  Web    `json:"web,omitempty"`
}
type Extattr struct {
	Attrs []Attrs `json:"attrs"`
}
type WechatChannels struct {
	Nickname string `json:"nickname"`
}
type Miniprogram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Title    string `json:"title"`
}
type ExternalAttr struct {
	Type        int         `json:"type"`
	Name        string      `json:"name"`
	Text        Text        `json:"text,omitempty"`
	Web         Web         `json:"web,omitempty"`
	Miniprogram Miniprogram `json:"miniprogram,omitempty"`
}
type ExternalProfile struct {
	ExternalCorpName string         `json:"external_corp_name"`
	WechatChannels   WechatChannels `json:"wechat_channels"`
	ExternalAttr     []ExternalAttr `json:"external_attr"`
}
