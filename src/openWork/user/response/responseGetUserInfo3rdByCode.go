package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetUserInfo3rdByCode struct {
	response.ResponseWork
	// CorpID 用户所属企业的corpid
	CorpID string `json:"corpid,omitempty"`
	// UserID 用户在企业内的UserID，如果该企业与第三方应用有授权关系时，返回明文UserId，否则返回密文UserId
	UserID string `json:"userid,omitempty"`
	// DeviceID 手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
	DeviceID string `json:"deviceid,omitempty"`
	// UserTicket 成员票据，最大为512字节。
	// scope为snsapi_userinfo或snsapi_privateinfo，且用户在应用可见范围之内时返回此参数。
	// 后续利用该参数可以获取用户信息或敏感信息，参见"第三方使用user_ticket获取成员详情"。
	UserTicket string `json:"user_ticket,omitempty"`
	// ExpiresIn user_ticket的有效时间（秒），随user_ticket一起返回
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// OpenUserID 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	OpenUserID string `json:"open_userid,omitempty"`
}
