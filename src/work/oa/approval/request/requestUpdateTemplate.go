package request

type TemplateName struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Title struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Placeholder struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Property struct {
	Control     string        `json:"control"`
	Id          string        `json:"id"`
	Title       []Title       `json:"title"`
	Placeholder []Placeholder `json:"placeholder"`
	Require     int           `json:"require"`
	UnPrint     int           `json:"un_print"`
}

type Control struct {
	Property Property `json:"property"`
	Config   struct {
	} `json:"config"`
}

type TemplateContent struct {
	Controls []Control `json:"controls"`
}

type RequestUpdateTemplate struct {
	TemplateId      string          `json:"template_id"`
	TemplateName    []TemplateName  `json:"template_name"`
	TemplateContent TemplateContent `json:"template_content"`
}
