package request

type RequestTagMarkTag struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}
