package request

type RequestDepartmentUpsert struct {
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
	ID       int    `json:"id"`
}
