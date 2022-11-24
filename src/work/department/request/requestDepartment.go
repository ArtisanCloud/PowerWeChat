package request

type RequestDepartmentInsert struct {
	Name     string `json:"name"`
	NameEn   string `json:"name_en,omitempty"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order,omitempty"`
	ID       int    `json:"id,omitempty"`
}

type RequestDepartmentUpdate struct {
	Name     string `json:"name,omitempty"`
	NameEn   string `json:"name_en,omitempty"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order,omitempty"`
	ID       int    `json:"id,omitempty"`
}
