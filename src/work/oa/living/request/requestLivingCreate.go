package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestLivingCreate struct {
	AnchorUserID         string         `json:"anchor_userid"`
	Theme                string         `json:"theme"`
	LivingStart          int64          `json:"living_start"`
	LivingDuration       int64          `json:"living_duration"`
	Description          string         `json:"description"`
	Type                 int            `json:"type"`
	AgentID              int            `json:"agentid"`
	RemindTime           int            `json:"remind_time"`
	ActivityCoverMediaID string         `json:"activity_cover_mediaid"`
	ActivityShareMediaID string         `json:"activity_share_mediaid"`
	ActivityDetail       *power.HashMap `json:"activity_detail"`
}
