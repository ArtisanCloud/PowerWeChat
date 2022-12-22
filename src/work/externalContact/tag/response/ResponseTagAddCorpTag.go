package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type Tag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

type CorpTagGroup struct {
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
	Tags       []*Tag `json:"tag"`
}

type ResponseTagAddCorpTag struct {
	response.ResponseWork

	TagGroups *CorpTagGroup `json:"tag_group"`
}
