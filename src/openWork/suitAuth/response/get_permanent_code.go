package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetPermanentCodeResponse struct {
	response.ResponseWork
	// PermanentCode 永久授权码，最长为512字节
	PermanentCode string `json:"permanent_code,omitempty"`
	// AuthCorpInfo 授权方企业信息
	AuthCorpInfo *AuthCorpInfo `json:"auth_corp_info,omitempty"`
	// AuthUserInfo 授权管理员的信息
	AuthUserInfo *AuthUserInfo `json:"auth_user_info,omitempty"`
	// V2
	RegisterCodeInfo `json:"register_code_info"`
	// State 安装应用时，扫码或者授权链接中带的state值。目前会返回state包含以下场景：扫带参二维码授权代开发模版
	State string `json:"state"`
}

type RegisterCodeInfo struct {
	// RegisterCode 注册码
	RegisterCode string `json:"register_code"`
	// TemplateId 推广包ID
	TemplateId string `json:"template_id"`
	// State 仅当获取注册码指定该字段时才返回
	State string `json:"state"`
}
type DealerCorpInfo struct {
	CorpId   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

// AuthCorpInfo 授权方企业信息
type AuthCorpInfo struct {
	// CorpID 授权方企业ID
	CorpID string `json:"corpid,omitempty"`
	// CorpName 授权方企业名
	CorpName string `json:"corp_name,omitempty"`
}

// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
type AuthInfo struct {
	// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
	Agent []Agent `json:"agent,omitempty"`
}

// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
type Agent struct {
	// AgentID 授权方应用id
	AgentID int `json:"agentid,omitempty"`
	// Name 授权方应用名字
	Name string `json:"name,omitempty"`
	// SquareLogoURL 授权方应用方形头像
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	// RoundLogoURL 授权方应用圆形头像
	RoundLogoURL string `json:"round_logo_url,omitempty"`
	// AppID 旧的多应用套件中的对应应用id，新开发者请忽略
	AppID uint64 `json:"appid,omitempty"`
	// Privilege 应用对应的权限
	Privilege Privilege `json:"privilege,omitempty"`
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
	// open_userid 授权管理员的open_userid，可能为空
	OpenUserID string `json:"open_userid,omitempty"`
	// Name 授权管理员的name，可能为空（内部管理员一定有，不可更改）
	Name string `json:"name,omitempty"`
	// Avatar 授权管理员的头像url
	Avatar string `json:"avatar,omitempty"`
}
