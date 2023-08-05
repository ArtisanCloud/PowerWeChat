package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type Invitor struct {
	UserID string `json:"userid"`
}

type Member struct {
	UserID        string   `json:"userid"`
	Type          int      `json:"type"`
	JoinTime      int      `json:"join_time"`
	JoinScene     int      `json:"join_scene"`
	State         string   `json:"state"`
	Invitor       *Invitor `json:"invitor,omitempty"`
	GroupNickname string   `json:"group_nickname"`
	Name          string   `json:"name"`
	UnionID       string   `json:"unionid,omitempty"`
}

type Admin struct {
	UserID string `json:"userid"`
}

type GroupChat struct {
	ChatID     string    `json:"chat_id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	CreateTime int       `json:"create_time"`
	Notice     string    `json:"notice"`
	MemberList []*Member `json:"member_list"`
	AdminList  []*Admin  `json:"admin_list"`
}

type ResponseGroupChatGet struct {
	response.ResponseWork
	GroupChat *GroupChat `json:"group_chat"`
}
