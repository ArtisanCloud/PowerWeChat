package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestUpgradeService struct {
	OpenKFID       string          `json:"open_kfid"`
	ExternalUserID string          `json:"external_userid"`
	Type           int             `json:"type"`
	Member         *power.StringMap `json:"member"`
	GroupChat      *power.StringMap `json:"groupchat"`
}
