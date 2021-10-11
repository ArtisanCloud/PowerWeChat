package request

type RequestAgentSet struct {
	AgentID            int    `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag"`
	LogoMediaID        string `json:"logo_mediaid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RedirectDomain     string `json:"redirect_domain"`
	IsReportEnter      int    `json:"isreportenter"`
	HomeUrl            string `json:"home_url"`
}