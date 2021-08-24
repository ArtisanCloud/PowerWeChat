package request

import "github.com/ArtisanCloud/go-libs/object"

type RequestUserDetail struct {
	UserID           string               `json:"userid"`
	Name             string               `json:"name"`
	Alias            string               `json:"alias"`
	Mobile           string               `json:"mobile"`
	Department       []int                `json:"department"`
	Order            []int                `json:"order"`
	Position         string               `json:"position"`
	Gender           string               `json:"gender"`
	Email            string               `json:"email"`
	IsLeaderInDept   []int                `json:"is_leader_in_dept"`
	Enable           int8                 `json:"enable"`
	AvatarMediaID    string               `json:"avatar_mediaid"`
	Telephone        string               `json:"telephone"`
	Address          string               `json:"address"`
	MainDepartment   int8                 `json:"main_department"`
	ExtAttr          *UserExtraAttribute  `json:"extattr"`
	ToInvite         bool                 `json:"to_invite"`
	ExternalPosition string               `json:"external_position"`
	ExternalProfile  *UserExternalProfile `json:"external_profile"`
}

type UserExtraAttribute struct {
	Attrs []*object.HashMap `json:"extattr"`
}

type UserExternalProfile struct {
	ExternalCorpName string            `json:"external_corp_name"`
	ExternalAttr     []*object.HashMap `json:"external_attr"`
}
