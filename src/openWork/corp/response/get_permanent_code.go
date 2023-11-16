package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetPermanentCodeResponse struct {
	response.ResponseWork
	// AccessToken 授权方（企业）access_token,最长为512字节
	AccessToken string `json:"access_token,omitempty"`
	// ExpiresIn 授权方（企业）access_token超时时间
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// PermanentCode 永久授权码，最长为512字节
	PermanentCode string `json:"permanent_code,omitempty"`
	// AuthCorpInfo 授权方企业信息
	AuthCorpInfo *AuthCorpInfo `json:"auth_corp_info,omitempty"`
	// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
	AuthInfo *AuthInfo `json:"auth_info,omitempty"`
	// AuthUserInfo 授权管理员的信息
	AuthUserInfo *AuthUserInfo `json:"auth_user_info,omitempty"`
}

// AuthCorpInfo 授权方企业信息
type AuthCorpInfo struct {
	// CorpID 授权方企业ID
	CorpID string `json:"corpid,omitempty"`
	// CorpName 授权方企业名
	CorpName string `json:"corp_name,omitempty"`
	// CorpType 授权方企业微信类型，认证号：verified, 注册号：unverified
	CorpType string `json:"corp_type,omitempty"`
	// CorpSquareLogoURL 授权方企业微信方形头像
	CorpSquareLogoURL string `json:"corp_square_logo_url,omitempty"`
	// CorpUserMax 授权方企业微信用户规模
	CorpUserMax int `json:"corp_user_max,omitempty"`
	// CorpFullName 所绑定的企业微信主体名称(仅认证过的企业有)
	CorpFullName string `json:"corp_full_name,omitempty"`
	// SubjectType 企业类型，1. 企业; 2. 政府以及事业单位; 3. 其他组织, 4.团队号
	SubjectType int `json:"subject_type,omitempty"`
	// VerfiedEndTime 认证到期时间
	VerfiedEndTime int64 `json:"verfied_end_time,omitempty"`
	// CorpWxqrcode 授权企业在微工作台（原企业号）的二维码，可用于关注微工作台
	CorpWxqrcode string `json:"corp_wxqrcode,omitempty"`
	// CorpScale 企业规模。当企业未设置该属性时，值为空
	CorpScale string `json:"corp_scale,omitempty"`
	// CorpIndustry 企业所属行业。当企业未设置该属性时，值为空
	CorpIndustry string `json:"corp_industry,omitempty"`
	// CorpSubIndustry 企业所属子行业。当企业未设置该属性时，值为空
	CorpSubIndustry string `json:"corp_sub_industry,omitempty"`
	// Location 企业所在地信息, 为空时表示未知
	Location string `json:"location,omitempty"`
}

// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
type AuthInfo struct {
	// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
	Agent []Agent `json:"agent,omitempty"`
}

// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
type Agent struct {
	// AgentID 授权方应用id
	AgentID int `json:"agent_id,omitempty"`
	// Name 授权方应用名字
	Name string `json:"name,omitempty"`
	// SquareLogoURL 授权方应用方形头像
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	// RoundLogoURL 授权方应用圆形头像
	RoundLogoURL string `json:"round_logo_url,omitempty"`
	// AppID 旧的多应用套件中的对应应用id，新开发者请忽略
	AppID uint64 `json:"appid,omitempty"`
	// Privilege 应用对应的权限
}

// Privilege 应用对应的权限
type Privilege struct {
	// AllowParty 应用可见范围（部门）
	AllowParty []uint64 `json:"allow_party,omitempty"`
	// AllowTag 应用可见范围（标签）
	AllowTag []uint64 `json:"allow_tag,omitempty"`
	// AllowUser 应用可见范围（用户）
	AllowUser []string `json:"allow_user,omitempty"`
	// ExtraParty 额外通讯录（部门）
	ExtraParty []uint64 `json:"extra_party,omitempty"`
	// ExtraUser 额外通讯录（成员）
	ExtraUser []string `json:"extra_user,omitempty"`
	// ExtraTag 额外通讯录（标签）
	ExtraTag []uint64 `json:"extra_tag,omitempty"`
	// Level 权限等级。
	// 1:通讯录基本信息只读
	// 2:通讯录全部信息只读
	// 3:通讯录全部信息读写
	// 4:单个基本信息只读
	// 5:通讯录全部信息只写
	Level int `json:"level,omitempty"`
}

// AuthUserInfo 授权管理员的信息
type AuthUserInfo struct {
	// UserID 授权管理员的userid，可能为空（内部管理员一定有，不可更改）
	UserID string `json:"userid,omitempty"`
	// Name 授权管理员的name，可能为空（内部管理员一定有，不可更改）
	Name string `json:"name,omitempty"`
	// Avatar 授权管理员的头像url
	Avatar string `json:"avatar,omitempty"`
}
