package request

type URLSchemeGenerate struct {
	JumpWxa        *JumpWxa `json:"jump_wxa"`
	IsExpire       bool     `json:"is_expire"`
	ExpireType     int      `json:"expire_type"`
	ExpireTime     int      `json:"expire_time,omitempty"`
	ExpireInterval int      `json:"expire_interval,omitempty"`
}

type JumpWxa struct {
	Path       string `json:"path,omitempty"`
	Query      string `json:"query,omitempty"`
	EnvVersion string `json:"env_version,omitempty"`
}
