package request

type RequestUserDetail struct {
	Userid           string          `json:"userid"`
	Name             string          `json:"name,omitempty"`
	Alias            string          `json:"alias,omitempty"`
	Mobile           string          `json:"mobile,omitempty"`
	Department       []int           `json:"department,omitempty"`
	Order            []int           `json:"order,omitempty"`
	Position         string          `json:"position,omitempty"`
	Gender           uint32          `json:"gender,omitempty"`
	Email            string          `json:"email,omitempty"`
	BizMail          string          `json:"biz_mail,omitempty"`
	IsLeaderInDept   []int           `json:"is_leader_in_dept,omitempty"`
	DirectLeader     []string        `json:"direct_leader,omitempty"`
	Enable           int             `json:"enable,omitempty"`
	AvatarMediaid    string          `json:"avatar_mediaid,omitempty"`
	Telephone        string          `json:"telephone,omitempty"`
	Address          string          `json:"address,omitempty"`
	MainDepartment   int             `json:"main_department,omitempty"`
	Extattr          Extattr         `json:"extattr,omitempty"`
	ToInvite         bool            `json:"to_invite,omitempty"`
	ExternalPosition string          `json:"external_position,omitempty"`
	ExternalProfile  ExternalProfile `json:"external_profile,omitempty"`
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
