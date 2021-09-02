package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAgentGet struct {
	*response.ResponseWork

	AgentID       int16          `json:"agentid"`         // "agentid": 1000005,
	Name          string         `json:"name"`            // "name": "HR助手",
	SquareLogoURL string         `json:"square_logo_url"` // "square_logo_url":  "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0",
	Description   string         `json:"description"`     // "description": "HR服务与员工自助平台",
	AllowUserInfo power.HashMap `json:"allow_userinfos"` // "allow_userinfos": {
	AllowParty    power.HashMap `json:"allow_partys"`    // "allow_partys": {
	AllowTags     power.HashMap `json:"allow_tags"`      // "allow_tags": {

	Close              int8   `json:"close"`
	RedirectDomain     string `json:"redirect_domai`
	ReportLocationFlag int8   `json:"report_location_flag"`
	IsReportEnter      int8   `json:"isreportente`
	HomeURL            string `json:"home_url"`
}
