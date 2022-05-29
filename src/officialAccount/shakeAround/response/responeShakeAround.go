package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type AccountRegisgerData struct {
	ApplyTime    int    `json:"apply_time"`
	AuditComment string `json:"audit_comment"`
	AuditStatus  int    `json:"audit_status"`
	AuditTime    int    `json:"audit_time"`
}

type ResponseShakeAroundAccountRegister struct {
	*response.ResponseOfficialAccount

	Data *AccountRegisgerData `json:"data"`
}

// ---------------------------------------------------------

type BeaconInfo struct {
	Distance float64 `json:"distance"`
	Major    int     `json:"major"`
	Minor    int     `json:"minor"`
	Uuid     string  `json:"uuid"`
}

type UserData struct {
	PageID     int         `json:"page_id "`
	BeaconInfo *BeaconInfo `json:"beacon_info"`
	Openid     string      `json:"openid"`
	PoiID      int         `json:" poi_id"`
}

type ResponseShakeAroundUser struct {
	*response.ResponseOfficialAccount

	Data *UserData `json:"data"`
}

// ---------------------------------------------------------

// ---------------------------------------------------------
