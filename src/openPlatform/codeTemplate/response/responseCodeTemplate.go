package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type Draft struct {
	CreateTime  int    `json:"create_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	DraftId     int    `json:"draft_id"`
}

type ResponseGetDrafts struct {
	response.ResponseOpenPlatform

	DraftList []*Draft `json:"draft_list"`
}

type Category struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	FirstId     int    `json:"first_id"`
	SecondId    int    `json:"second_id"`
}

type Template struct {
	CreateTime             int         `json:"create_time"`
	UserVersion            string      `json:"user_version"`
	UserDesc               string      `json:"user_desc"`
	TemplateId             int         `json:"template_id"`
	SourceMiniProgramAppID string      `json:"source_miniprogram_appid"`
	SourceMiniProgram      string      `json:"source_miniprogram"`
	Developer              string      `json:"developer"`
	TemplateType           int         `json:"template_type"`
	CategoryList           []*Category `json:"category_list"`
	AuditScene             int         `json:"audit_scene,omitempty"`
	AuditStatus            int         `json:"audit_status,omitempty"`
	Reason                 string      `json:"reason,omitempty"`
}

type ResponseList struct {
	response.ResponseOpenPlatform

	TemplateList []*Template `json:"template_list"`
}
