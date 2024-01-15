package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseGetActiveInfoByCode struct {
	response.ResponseWork
	ActiveInfo *model.ActiveInfo `json:"active_info,omitempty"`
}

type ResponseBatchGetActiveInfoByCode struct {
	response.ResponseWork
	ActiveInfoList []model.ActiveInfo `json:"active_info_list,omitempty"`
}

type ResponseGetActiveInfoByUser struct {
	response.ResponseWork
	// ActiveStatus 账号激活状态。0：未激活、 1：已激活
	ActiveStatus   int                `json:"active_status,omitempty"`
	ActiveInfoList []model.ActiveInfo `json:"active_info_list,omitempty"`
}
