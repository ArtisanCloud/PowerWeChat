package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseAgentGet struct {
	response.ResponseWork
	AgentID            int                         `json:"agentid"`         // "agentid": 1000005,
	Name               string                      `json:"name"`            // "name": "HR助手",
	SquareLogoURL      string                      `json:"square_logo_url"` // "square_logo_url":  "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0",
	Description        string                      `json:"description"`     // "description": "HR服务与员工自助平台",
	AllowUserInfos     ResponseAgentAllowUserInfos `json:"allow_userinfos"` // "allow_userinfos": {
	AllowParty         ResponseAgentAllowParty     `json:"allow_partys"`    // "allow_partys": {
	AllowTags          ResponseAgentAllowTags      `json:"allow_tags"`      // "allow_tags": {
	Close              int8                        `json:"close"`
	RedirectDomain     string                      `json:"redirect_domain"`
	ReportLocationFlag int8                        `json:"report_location_flag"`
	IsReportEnter      int8                        `json:"isreportenter"`
	HomeURL            string                      `json:"home_url"`
}

type ResponseAgentAllowUserInfos struct {
	response.ResponseWork

	User []*ResponseAgentAllowUser `json:"user"`
}

type ResponseAgentAllowUser struct {
	response.ResponseWork

	UserID string `json:"userid"`
}

type ResponseAgentAllowParty struct {
	response.ResponseWork

	PartyID []int `json:"partyid"`
}

type ResponseAgentAllowTags struct {
	TagID []int `json:"tagid"`
}
