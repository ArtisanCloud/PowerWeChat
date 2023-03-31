package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/goods/request"
)

type ResponseProductAdd struct {
	response.ResponseOfficialAccount

	StatusTicket string `json:"status_ticket"`
}

// ----------------------------------------

type ResponseProductStatus struct {
	response.ResponseOfficialAccount

	Result struct {
		SuccCnt  int    `json:"succ_cnt"`
		FailCnt  int    `json:"fail_cnt"`
		TotalCnt int    `json:"total_cnt"`
		Progress string `json:"progress"`
		Statuses []struct {
			Pid        string `json:"pid"`
			Ret        int    `json:"ret"`
			ErrMsg     string `json:"err_msg"`
			ErrMsgZhCn string `json:"err_msg_zh_cn"`
		} `json:"statuses"`
	} `json:"result"`
}

// ----------------------------------------

type ResponseProductGet struct {
	response.ResponseOfficialAccount

	Product *request.Product `json:"product"`
}
