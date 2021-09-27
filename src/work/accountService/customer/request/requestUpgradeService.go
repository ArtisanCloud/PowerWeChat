package request

type RequestUpgradeService struct {
	OpenKFID       string                          `json:"open_kfid"`
	ExternalUserID string                          `json:"external_userid"`
	Type           int                             `json:"type"`
	Member         *RequestUpgradeServiceMember    `json:"member,omitempty"`
	GroupChat      *RequestUpgradeServiceGroupChat `json:"groupchat,omitempty"`
}

type RequestUpgradeServiceMember struct {
	UserID  string `json:"userid"`
	Wording string `json:"wording"`
}

type RequestUpgradeServiceGroupChat struct {
	ChatID  string `json:"chat_id"`
	Wording string `json:"wording"`
}
