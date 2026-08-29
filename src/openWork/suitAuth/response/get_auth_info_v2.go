package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type GetAuthInfoV2 struct {
	response.ResponseWork
	// AuthCorpInfo 授权方企业信息
	AuthCorpInfo *AuthCorpInfoV2 `json:"auth_corp_info,omitempty"`
	// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的
	AuthInfo *AuthInfoV2 `json:"auth_info,omitempty"`
	// DealerCorpInfo 代理服务商企业信息
	DealerCorpInfo *DealerCorpInfoV2 `json:"dealer_corp_info,omitempty"`
}

// AuthCorpInfoV2 授权方企业信息
type AuthCorpInfoV2 struct {
	// CorpID 授权方企业微信 id
	CorpID string `json:"corpid,omitempty"`
	// CorpName 授权方企业名称
	CorpName string `json:"corp_name,omitempty"`
	// CorpType 授权方企业类型，认证号：verified, 注册号：unverified
	CorpType string `json:"corp_type,omitempty"`
	// CorpSquareLogoURL 授权方企业方形头像
	CorpSquareLogoURL string `json:"corp_square_logo_url,omitempty"`
	// CorpUserMax 授权方企业用户规模
	CorpUserMax int `json:"corp_user_max,omitempty"`
	// CorpFullName 授权方企业的主体名称 (仅认证或验证过的企业有)，即企业全称
	CorpFullName string `json:"corp_full_name,omitempty"`
	// SubjectType 企业类型，1. 企业; 2. 政府以及事业单位; 3. 其他组织，4.团队号
	SubjectType int `json:"subject_type,omitempty"`
	// VerifiedEndTime 认证到期时间
	VerifiedEndTime int64 `json:"verified_end_time,omitempty"`
	// CorpScale 企业规模。当企业未设置该属性时，值为空
	CorpScale string `json:"corp_scale,omitempty"`
	// CorpIndustry 企业所属行业。当企业未设置该属性时，值为空
	CorpIndustry string `json:"corp_industry,omitempty"`
	// CorpSubIndustry 企业所属子行业。当企业未设置该属性时，值为空
	CorpSubIndustry string `json:"corp_sub_industry,omitempty"`
	// CorpExName 企业其他认证的名称，仅认证企业才有
	CorpExName *CorpExName `json:"corp_ex_name,omitempty"`
}

// CorpExName 企业其他认证的名称
type CorpExName struct {
	// NameList 企业其他认证的企业简称列表（不包括 corp_name）
	NameList []string `json:"name_list,omitempty"`
}

// AuthInfoV2 授权信息
type AuthInfoV2 struct {
	// Agent 授权的应用信息，注意是一个数组
	Agent []AgentV2 `json:"agent,omitempty"`
}

// AgentV2 授权的应用信息
type AgentV2 struct {
	// AgentID 授权方应用 id
	AgentID int `json:"agentid,omitempty"`
	// Name 授权方应用名字
	Name string `json:"name,omitempty"`
	// SquareLogoURL 授权方应用方形头像
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	// RoundLogoURL 授权方应用圆形头像
	RoundLogoURL string `json:"round_logo_url,omitempty"`
	// AppID 旧的多应用套件中的对应应用 id，新开发者请忽略
	AppID uint64 `json:"appid,omitempty"`
	// AuthMode 授权模式，0 为管理员授权；1 为成员授权
	AuthMode int `json:"auth_mode,omitempty"`
	// IsCustomizedApp 是否为代开发自建应用
	IsCustomizedApp bool `json:"is_customized_app,omitempty"`
	// Privilege 应用对应的权限
	Privilege *PrivilegeV2 `json:"privilege,omitempty"`
	// SharedFrom 共享了应用的企业信息，仅当由企业互联或者上下游共享应用触发的安装时才返回
	SharedFrom *SharedFrom `json:"shared_from,omitempty"`
}

// PrivilegeV2 应用对应的权限
type PrivilegeV2 struct {
	// AllowParty 应用可见范围（部门）
	AllowParty []uint64 `json:"allow_party,omitempty"`
	// AllowTag 应用可见范围（标签）
	AllowTag []uint64 `json:"allow_tag,omitempty"`
	// AllowUser 应用可见范围（成员）
	AllowUser []string `json:"allow_user,omitempty"`
	// ExtraParty 额外通讯录（部门）
	ExtraParty []uint64 `json:"extra_party,omitempty"`
	// ExtraUser 额外通讯录（成员）
	ExtraUser []string `json:"extra_user,omitempty"`
	// ExtraTag 额外通讯录（标签）
	ExtraTag []uint64 `json:"extra_tag,omitempty"`
	// Level 权限等级。
	// 1:通讯录基本信息只读
	// 2:通讯录全部信息只读（已废弃）
	// 3:通讯录全部信息读写
	// 4:单个基本信息只读
	// 5:通讯录全部信息只写（已废弃）
	// 0:为代开发应用，该值无意义，固定为 0
	Level int `json:"level,omitempty"`
}

// SharedFrom 共享了应用的企业信息
type SharedFrom struct {
	// CorpID 共享了应用的企业 corpid，仅当企业互联或者上下游共享应用触发的安装时才返回
	CorpID string `json:"corpid,omitempty"`
	// ShareType 共享了途径，0 表示企业互联，1 表示上下游
	ShareType int `json:"share_type,omitempty"`
}

// DealerCorpInfoV2 代理服务商企业信息
type DealerCorpInfoV2 struct {
	// CorpID 代理服务商企业微信 id
	CorpID string `json:"corpid,omitempty"`
	// CorpName 代理服务商企业微信名称
	CorpName string `json:"corp_name,omitempty"`
}
