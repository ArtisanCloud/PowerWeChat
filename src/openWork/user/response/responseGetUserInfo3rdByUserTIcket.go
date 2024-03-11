package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetUserInfo3rdByUserTicket struct {
	response.ResponseWork
	// CorpID 用户所属企业的corpid
	CorpID string `json:"corpid,omitempty"`
	// UserID 用户在企业内的UserID，如果该企业与第三方应用有授权关系时，返回明文UserId，否则返回密文UserId
	UserID string `json:"userid,omitempty"`
	// Name 成员姓名
	Name string `json:"name,omitempty"`
	// Mobile 成员手机号，仅在用户同意snsapi_privateinfo授权时返回
	Mobile string `json:"mobile,omitempty"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性
	Gender int `json:"gender,omitempty"`
	// Email 成员邮箱，仅在用户同意snsapi_privateinfo授权时返回
	Email string `json:"email,omitempty"`
	// Avatar 头像url。注：如果要获取小图将url最后的"/0"改成"/100"即可。仅在用户同意snsapi_privateinfo授权时返回
	Avatar string `json:"avatar,omitempty"`
	// Qrcode 员工个人二维码（扫描可添加为外部联系人），仅在用户同意snsapi_privateinfo授权时返回
	Qrcode string `json:"qrcode,omitempty"`
}
