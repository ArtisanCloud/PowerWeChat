package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"

type RequestSubmitOrderJob struct {
	// JobID 任务id
	JobID string `json:"jobid,omitempty"`
	// BuyerUserID 下单人，服务商企业内成员userid。该userid必须登录过企业微信，并且企业微信已绑定微信，且必须为服务商企业内具有“购买接口许可”权限的管理员。
	BuyerUserID string `json:"buyer_userid,omitempty"`
	// AccountDuration 账号购买时长
	AccountDuration *model.AccountDuration `json:"account_duration,omitempty"`
}
