package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type UserSummary struct {
	RefDate    string `json:"ref_date"`
	UserSource int    `json:"user_source"`
	NewUser    int    `json:"new_user"`
	CancelUser int    `json:"cancel_user"`
}

type ResponseDataCubeUserSummary struct {
	*response.ResponseOfficialAccount

	List []*UserSummary `json:"list"`
}

// ---------------------------------------------------------

type UserCumulate struct {
	RefDate      string `json:"ref_date"`
	CumulateUser int    `json:"cumulate_user"`
}

type ResponseDataCubeUserCumulate struct {
	*response.ResponseOfficialAccount

	List []*UserCumulate `json:"list"`
}
