package request

type Item struct {
	Address     string `json:"address"`
	Tag         string `json:"tag"`
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	FirstID     int    `json:"first_id"`
	SecondID    int    `json:"second_id"`
	Title       string `json:"title"`
	ThirdClass  string `json:"third_class,omitempty"`
	ThirdID     int    `json:"third_id,omitempty"`
}

type PreviewInfo struct {
	VideoIDList []string `json:"video_id_list"`
	PicIDList   []string `json:"pic_id_list"`
}

type UGCDeclare struct {
	Scene        []int  `json:"scene"`
	Method       []int  `json:"method"`
	HasAuditTeam int    `json:"has_audit_team"`
	AuditDesc    string `json:"audit_desc"`
}

type RequestSubmitAudit struct {
	ItemList      []*Item      `json:"item_list"`
	FeedbackInfo  string       `json:"feedback_info"`
	FeedbackStuff string       `json:"feedback_stuff"`
	PreviewInfo   *PreviewInfo `json:"preview_info"`
	VersionDesc   string       `json:"version_desc"`
	UGCDeclare    *UGCDeclare  `json:"ugc_declare"`
}
