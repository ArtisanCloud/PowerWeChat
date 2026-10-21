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

// RequestDepartmentUpdatePartial 用于仅更新显式设置的部门字段。
// ID 为必填；其余字段为 nil 时不修改，非 nil 时发送指向的值，包括空字符串和零值。
// 例如，NameEn: object.String("") 表示清空英文名称。
type RequestDepartmentUpdatePartial struct {
	ID       int     `json:"id"`
	Name     *string `json:"name,omitempty"`
	NameEn   *string `json:"name_en,omitempty"`
	ParentID *int    `json:"parentid,omitempty"`
	Order    *int    `json:"order,omitempty"`
}
