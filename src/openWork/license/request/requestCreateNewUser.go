package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"

type RequestCreateNewOrder struct {
	CorpID string `json:"corpid,omitempty"`
	// BuyerUserID 下单人。服务商企业内成员userid。该userid必须登录过企业微信，并且企业微信已绑定微信，且必须为服务商企业内具有“购买接口许可”权限的管理员。最终也支持由其他人支付
	BuyerUserID string `json:"buyer_userid,omitempty"`
	// AccountCount 账号个数详情，基础账号跟互通账号不能同时为0
	AccountCount *model.AccountCount `json:"account_count,omitempty"`
	// AccountDuration 账号购买时长。总购买时长为(months*31+days)天，最少购买1个月(31天)，最多购买60个月(1860天)。若企业为服务商测试企业，只支持购买1个月，不支持指定天购买
	AccountDuration *model.AccountDuration `json:"account_duration,omitempty"`
}
