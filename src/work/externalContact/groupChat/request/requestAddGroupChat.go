package request

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"

type RequestAddGroupChat struct {
	StatusFilter int            `json:"status_filter"`
	OwnerFilter  *power.HashMap `json:"owner_filter"`
	Cursor       string         `json:"cursor"`
	Limit        int            `json:"limit"`
}
