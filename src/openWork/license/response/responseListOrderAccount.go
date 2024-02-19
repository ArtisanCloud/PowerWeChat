package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseListOrderAccount struct {
	response.ResponseWork
	NextCursor  string             `json:"next_cursor,omitempty"`
	HasMore     int                `json:"has_more,omitempty"`
	AccountList []model.ActiveInfo `json:"account_list,omitempty"`
}
