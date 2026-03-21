package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetUserInfo3rdByUserTicket struct {
	response.ResponseWork
	// CorpID 用户所属企业的corpid
	CorpID string `json:"corpid,omitempty"`
	// UserID 用户在企业内的UserID，如果该企业与第三方应用有授权关系时，返回明文UserId，否则返回密文UserId
	UserID string `json:"userid,omitempty"`
	// Name 成员姓名，此字段从2019年12月30日起，对新创建第三方应用不再返回真实name，使用userid代替name返回，2020年6月30日起，对所有历史第三方应用不再返回，第三方页面需要通过通讯录展示组件来展示名字
	Name string `json:"name,omitempty"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性。仅在用户同意snsapi_privateinfo授权时返回真实值，否则返回0.
	Gender int `json:"gender,omitempty"`
	// Avatar 头像url。仅在用户同意snsapi_privateinfo授权时返回真实头像，否则返回默认头像
	Avatar string `json:"avatar,omitempty"`
	// Qrcode 员工个人二维码（扫描可添加为外部联系人），仅在用户同意snsapi_privateinfo授权时返回
	Qrcode string `json:"qrcode,omitempty"`
}
