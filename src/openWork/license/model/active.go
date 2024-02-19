package model

type ActiveInfo struct {
	ActiveCode string `json:"active_code,omitempty"`
	UserID     string `json:"userid,omitempty"`
	CorpID     string `json:"corpid,omitempty"`
	// Type 账号类型：1:基础账号，2:互通账号
	Type int `json:"type,omitempty"`
	// Status 账号状态：
	// 1: 未绑定
	// 2: 已绑定且有效
	// 3: 已过期
	// 4: 待转移(企业开启自动激活时，成员离职或者被移出可见范围，第二天凌晨会更新为该状态)
	// 5: 已合并（激活码本身激活了userid，后续使用新的激活码重新激活了该userid，则该码变为已合并状态。若被合并时，该激活码未过期则合并后会重置expire_time为合并时间。若被合并时，激活码已过期则不重置expire_time。注：该状态的激活码是已经失效的，不能重新用于激活或者继承。）
	// 6: 已分配给下游
	Status int `json:"status,omitempty"`
	// CreateTime 建时间，订单支付成功后立即创建。
	CreateTime int64 `json:"create_time,omitempty"`
	// ActiveTime 首次激活绑定用户的时间，未激活则不返回该字段
	ActiveTime int64 `json:"active_time,omitempty"`
	// ExpireTime 过期时间。为首次激活绑定的时间加上购买时长。未激活则不返回该字段
	ExpireTime int64 `json:"expire_time,omitempty"`
	// MergeInfo 合并信息，合并的激活码或者被合并的激活码才返回该字段
	MergeInfo *ActiveMergeInfo `json:"merge_info,omitempty"`
	// ShareInfo 分配信息，当激活码在上下游/企业互联场景下，从上游分配给下游时，获取上游或者下游企业该激活码详情时返回
	ShareInfo *ActiveShareInfo `json:"share_info,omitempty"`
}

// ActiveMergeInfo 合并信息，合并的激活码或者被合并的激活码才返回该字段
type ActiveMergeInfo struct {
	// ToActiveCode 该激活码合并到的新激活码信息
	ToActiveCode string `json:"to_active_code,omitempty"`
	// FromActiveCode 激活码激活userid时，若userid原来已经绑定了一个激活码，则会返回该字段
	FromActiveCode string `json:"from_active_code,omitempty"`
}

// ActiveShareInfo 分配信息，当激活码在上下游/企业互联场景下，从上游分配给下游时，获取上游或者下游企业该激活码详情时返回
type ActiveShareInfo struct {
	// ToCorpID 下游企业corpid。当激活码通过上游分配给下游时，获取上游企业该激活码详情时返回该字段，表示被分配给了哪个下游企业
	ToCorpID string `json:"to_corpid,omitempty"`
	// FromCorpID 上游企业corpid。当激活码通过上游分配给下游时，获取下游企业该激活码详情时返回该字段，表示从哪个上游企业分配过来
	FromCorpID string `json:"from_corpid,omitempty"`
}
