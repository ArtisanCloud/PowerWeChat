package request

type CloudBase struct {
	Env    string `json:"env"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	Query  string `json:"query"`
}

type URLLinkGenerate struct {
	Path           string     `json:"path"`
	Query          string     `json:"query"`
	ExpireType     int        `json:"expire_type"`
	ExpireInterval int        `json:"expire_interval"`
	ExpireTime     int        `json:"expire_time"`
	EnvVersion     string     `json:"env_version"`
	CloudBase      *CloudBase `json:"cloud_base"`
}
