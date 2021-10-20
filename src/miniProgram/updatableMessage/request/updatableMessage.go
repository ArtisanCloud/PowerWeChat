package request

type RequestSetUpdatableMsg struct {
	ActivityID   string        `json:"activity_id"`
	TargetState  int8          `json:"target_state"`
	TemplateInfo *TemplateInfo `json:"template_info"`
}

type TemplateInfo struct {
	ParameterList []*ParameterListItem `json:"parameter_list"`
}
type ParameterListItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
