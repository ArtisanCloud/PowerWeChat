package request

type SetSessionInfoRequest struct {
	// PreAuthCode 预授权码
	PreAuthCode string `json:"pre_auth_code,omitempty"`
	// SessionInfo 本次授权过程中需要用到的会话信息
}

// SessionInfo 本次授权过程中需要用到的会话信息
type SessionInfo struct {
	// AppID 允许进行授权的应用id，如1、2、3， 不填或者填空数组都表示允许授权套件内所有应用（仅旧的多应用套件可传此参数，新开发者可忽略）
	AppID []uint64 `json:"appid,omitempty"`
	// AuthType 授权类型：0 正式授权， 1 测试授权。 默认值为0。注意，请确保应用在正式发布后的授权类型为“正式授权”
	AuthType int `json:"auth_type,omitempty"`
}
